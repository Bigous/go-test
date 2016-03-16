package utils

import "testing";

func TestReverso(this *testing.T) {
	var casos = []struct { entrada, esperado string } {
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	};
	var retorno string;
	for _, c := range casos {
		retorno = Reverso(c.entrada);
		if (retorno != c.esperado) {
			this.Errorf("Reverso(%q) == %q, esperado %q", c.entrada, retorno, c.esperado);
		}
	}
}