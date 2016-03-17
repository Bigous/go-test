package utils

import "reflect"
import "fmt"

// ReverteString retorna uma string com os caracteres invertido em relação a string passada.
func ReverteString(s string) string {
	var (
		r       []rune
		i, j, t int
	)
	r = []rune(s)
	t = len(r)
	for i, j = 0, t-1; i < t/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ReverteArray reverte um array de qualquer tipo.
func ReverteArray(array Any) Any {
	varray := reflect.ValueOf(array)
	tipo := reflect.TypeOf(array)
	kind := tipo.Kind()
	if kind == reflect.String {
		return ReverteString(array.(string))
	}
	if tipo.Kind() != reflect.Array && kind != reflect.String && kind != reflect.Slice {
		panic(fmt.Errorf("Not an array %q - %q", array, kind))
	}
	t := varray.Len()
	ret := make(AnyArray, t)
	for i, j := 0, t-1; j >= 0; i, j = i+1, j-1 {
		ret[i] = varray.Index(j).Interface()
	}
	return ret
}
