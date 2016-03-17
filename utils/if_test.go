package utils

import (
	"testing"
)

func TestIf(t *testing.T) {
	casos := []struct {
		condicao bool
		pTrue    interface{}
		pFalse   interface{}
		esperado interface{}
	}{
		{true, "ae", "uh", "ae"},
		{false, "uh", "ae", "ae"},
	}

	for _, c := range casos {
		v := If(c.condicao, c.pTrue, c.pFalse)
		if v != c.esperado {
			t.Errorf("If(%t, %q, %q) = %q should be %q", c.condicao, c.pTrue, c.pFalse, v, c.esperado)
		}
	}

}
