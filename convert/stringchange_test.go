package convert_test

import (
	"testing"

	"github.com/SVz777/tk/convert"
)

func TestNewStringChange(t *testing.T) {
	tests := []struct {
		s    string
		real []string
	}{
		{
			s:    "abcd",
			real: []string{"abcd"},
		},
		{
			s:    "a bc d",
			real: []string{"a", "bc", "d"},
		},
		{
			s:    "a bc-d",
			real: []string{"a", "bc", "d"},
		},
		{
			s:    "a_bc-d",
			real: []string{"a", "bc", "d"},
		},
		{
			s:    "A_Bc-d",
			real: []string{"a", "bc", "d"},
		},
		{
			s:    "aBCd",
			real: []string{"a", "b", "cd"},
		},
		{
			s:    "a_Bc-d",
			real: []string{"a", "bc", "d"},
		},
		{
			s:    "a_bC-d",
			real: []string{"a", "b", "c", "d"},
		},
	}
NEXT:
	for _, test := range tests {
		sc := convert.NewStringChange(test.s)
		lsc := len(sc.Words)
		ltr := len(test.real)
		if lsc != ltr {
			t.Error(test.s, lsc, ltr)
			continue
		}
		for i := 0; i < lsc; i++ {
			if sc.Words[i] != test.real[i] {
				t.Error(test.s, sc.Words, test.real)
				continue NEXT
			}
		}
		t.Log(test.s, sc.Words, test.real)
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		s    string
		real string
	}{
		{
			s:    "abcd",
			real: "abcd",
		},
		{
			s:    "a bc d",
			real: "aBcD",
		},
		{
			s:    "a bc-d",
			real: "aBcD",
		},
		{
			s:    "a_bc-d",
			real: "aBcD",
		},
		{
			s:    "A_Bc-d",
			real: "aBcD",
		},
		{
			s:    "aBCd",
			real: "aBCd",
		},
		{
			s:    "A_Bc-d",
			real: "aBcD",
		},
		{
			s:    "A_bC-d",
			real: "aBCD",
		},
	}
	for _, test := range tests {
		z := convert.ToCamelCase(test.s)
		if z != test.real {
			t.Error(test.s, z, test.real)
			continue
		}
		t.Log(test.s, z, test.real)
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		s    string
		real string
	}{
		{
			s:    "abcd",
			real: "abcd",
		},
		{
			s:    "a bc d",
			real: "a_bc_d",
		},
		{
			s:    "a bc-d",
			real: "a_bc_d",
		},
		{
			s:    "a_bc-d",
			real: "a_bc_d",
		},
		{
			s:    "A_Bc-d",
			real: "a_bc_d",
		},
		{
			s:    "aBCd",
			real: "a_b_cd",
		},
		{
			s:    "a_Bc-d",
			real: "a_bc_d",
		},
		{
			s:    "a_bC-d",
			real: "a_b_c_d",
		},
	}
	for _, test := range tests {
		z := convert.ToSnakeCase(test.s)
		if z != test.real {
			t.Error(test.s, z, test.real)
			continue
		}
		t.Log(test.s, z, test.real)
	}
}

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		s    string
		real string
	}{
		{
			s:    "abcd",
			real: "abcd",
		},
		{
			s:    "a bc d",
			real: "a-bc-d",
		},
		{
			s:    "a bc-d",
			real: "a-bc-d",
		},
		{
			s:    "a_bc-d",
			real: "a-bc-d",
		},
		{
			s:    "A_Bc-d",
			real: "a-bc-d",
		},
		{
			s:    "aBCd",
			real: "a-b-cd",
		},
		{
			s:    "a_Bc-d",
			real: "a-bc-d",
		},
		{
			s:    "a_bC-d",
			real: "a-b-c-d",
		},
	}
	for _, test := range tests {
		z := convert.ToKebabCase(test.s)
		if z != test.real {
			t.Error(test.s, z, test.real)
			continue
		}
		t.Log(test.s, z, test.real)
	}
}
