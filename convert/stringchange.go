package convert

import (
	"strings"
	"unicode"
)

type StringCase = uint8

const (
	TypeCamelCase StringCase = iota
	TypeSnakeCase
	TypeKebabCase
	TypeMaxCase
)

type StringChange struct {
	Words []string
}

func NewStringChange(s string) *StringChange {
	sc := &StringChange{}
	sc.Reset(s)
	return sc
}

func ToCamelCase(s string) string {
	sc := NewStringChange(s)
	return sc.To(TypeCamelCase)
}
func ToSnakeCase(s string) string {
	sc := NewStringChange(s)
	return sc.To(TypeSnakeCase)
}
func ToKebabCase(s string) string {
	sc := NewStringChange(s)
	return sc.To(TypeKebabCase)
}

func (sc *StringChange) Reset(s string) {
	var words []string
	sb := strings.Builder{}
	_ = sb
	for _, i := range s {
		if i == '_' || i == '-' || i == ' ' {
			words = append(words, sb.String())
			sb.Reset()
			continue
		}
		if 'A' <= i && i <= 'Z' {
			if sb.Len() != 0 {
				words = append(words, sb.String())
				sb.Reset()
			}
			sb.WriteRune(unicode.ToLower(i))
			continue
		}
		sb.WriteRune(i)
	}
	if sb.Len() != 0 {
		words = append(words, sb.String())
	}

	sc.Words = words
}

func (sc *StringChange) To(_case StringCase) string {
	if len(sc.Words) == 0 || (_case < 0 || _case >= TypeMaxCase) {
		return ""
	}

	switch _case {
	case TypeCamelCase:
		sb := strings.Builder{}
		sb.WriteString(sc.Words[0])
		for _, word := range sc.Words[1:] {
			sb.WriteString(strings.Title(word))
		}
		return sb.String()
	case TypeSnakeCase:
		return strings.Join(sc.Words, "_")
	case TypeKebabCase:
		return strings.Join(sc.Words, "-")
	}

	return ""
}
