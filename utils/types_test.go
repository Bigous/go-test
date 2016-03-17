package utils

import "testing"

func TestAny(t *testing.T) {
	var (
		inteiro Any
		texto   Any
	)
	inteiro = 234234
	texto = "texto"
	if inteiro != 234234 {
		t.Errorf("Inteiro devia ser %q mas é %q", 234234, inteiro)
	}
	if texto != "texto" {
		t.Errorf("Texto devia ser %q mas é %q", "texto", texto)
	}
}

func TestAnyArray(t *testing.T) {
	var arr AnyArray
	arr = make(AnyArray, 3)
	arr[0] = 234234
	arr[1] = "teste"
	arr[2] = 'a'
	if arr[0] != 234234 || arr[1] != "teste" || arr[2] != 'a' {
		t.Errorf("Erro no array %q", arr)
	}
}

func TestMapStringAny(t *testing.T) {
	var json = MapStringAny{
		"inteiro":  231231,
		"texto":    "teste de escrita",
		"caracter": 'a',
		"objeto": MapStringAny{
			"inteiro":  231231,
			"texto":    "teste de escrita",
			"caracter": 'a',
		},
	}
	if json == nil {
		t.Errorf("Json %q", json)
	}
}

func TestJSON(t *testing.T) {
	var json = JSON{
		"inteiro":  231231,
		"texto":    "teste de escrita",
		"caracter": 'a',
		"objeto": JSON{
			"inteiro":  231231,
			"texto":    "teste de escrita",
			"caracter": 'a',
		},
	}
	if json == nil {
		t.Errorf("Json %q", json)
	}
}
