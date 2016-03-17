package utils

import (
	"reflect"
	"testing"
)

func TestReverteString(this *testing.T) {
	var casos = []struct{ entrada, esperado string }{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	var retorno string
	for _, c := range casos {
		retorno = ReverteString(c.entrada)
		if retorno != c.esperado {
			this.Errorf("ReverteString(%q) == %q, esperado %q", c.entrada, retorno, c.esperado)
		}
	}
}

func testEq(a, b Any, t *testing.T) bool {
	if a == nil && b == nil {
		return true
	}
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	if a == nil || b == nil || va.Kind() != vb.Kind() || va.Len() != vb.Len() {
		return false
	}
	for i := 0; i < va.Len(); i++ {
		if va.Index(i).Interface() != vb.Index(i).Interface() {
			return false
		}
	}
	return true
}

func TestReverteArray(this *testing.T) {
	casos := JSON{
		"strings":          AnyArray{"Hello, world", "dlrow ,olleH"},
		"runas":            AnyArray{[]rune{'a', 'b', 'c'}, []rune{'c', 'b', 'a'}},
		"inteiros":         AnyArray{[]int{1, 2, 3}, []int{3, 2, 1}},
		"array de strings": AnyArray{[]string{"teste", "de", "sorvete"}, []string{"sorvete", "de", "teste"}},
	}
	for _, c := range casos {
		entrada := reflect.ValueOf(c).Index(0).Interface()
		esperado := reflect.ValueOf(c).Index(1).Interface()
		retorno := ReverteArray(entrada)
		if !testEq(retorno, esperado, this) {
			this.Errorf("ReverteArray(%q) == %q, esperado %q", entrada, retorno, esperado)
		}
	}
}
