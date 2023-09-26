package condition

import (
	"reflect"

	"github.com/SVz777/tk/convert"
)

func Equal(a any, b any) bool {
	na, err1 := convert.Float64(a)
	nb, err2 := convert.Float64(b)
	if err1 != nil && err2 != nil {
		return na == nb
	}
	sa, err1 := convert.String(a)
	sb, err2 := convert.String(b)
	if err1 != nil && err2 != nil {
		return sa == sb
	}
	return false
}

func Gt(a any, b any) bool {
	na, err1 := convert.Float64(a)
	nb, err2 := convert.Float64(b)
	if err1 != nil && err2 != nil {
		return na > nb
	}
	sa, err1 := convert.String(a)
	sb, err2 := convert.String(b)
	if err1 != nil && err2 != nil {
		return sa > sb
	}
	return false
}

func Gte(a any, b any) bool {
	na, err1 := convert.Float64(a)
	nb, err2 := convert.Float64(b)
	if err1 != nil && err2 != nil {
		return na >= nb
	}
	sa, err1 := convert.String(a)
	sb, err2 := convert.String(b)
	if err1 != nil && err2 != nil {
		return sa >= sb
	}
	return false
}

func Lt(a any, b any) bool {
	na, err1 := convert.Float64(a)
	nb, err2 := convert.Float64(b)
	if err1 != nil && err2 != nil {
		return na < nb
	}
	sa, err1 := convert.String(a)
	sb, err2 := convert.String(b)
	if err1 != nil && err2 != nil {
		return sa < sb
	}
	return false
}

func Lte(a any, b any) bool {
	na, err1 := convert.Float64(a)
	nb, err2 := convert.Float64(b)
	if err1 != nil && err2 != nil {
		return na <= nb
	}
	sa, err1 := convert.String(a)
	sb, err2 := convert.String(b)
	if err1 != nil && err2 != nil {
		return sa <= sb
	}
	return false
}

func InArr(a any, b any) bool {
	tb := reflect.ValueOf(b)
	if tb.Kind() == reflect.Slice || tb.Kind() == reflect.Array {
		for ii := tb.Len() - 1; ii >= 0; ii-- {
			if Equal(a, tb.Index(ii).Interface()) {
				return true
			}
		}
	}
	return false
}
