package utils

// If faz um if de linha. Caso a condição seja verdadeira, retorna pTrue, caso seja falsa, retorna pFalse.
func If(condition bool, pTrue interface{}, pFalse interface{}) interface{} {
	if condition {
		return pTrue
	}
	return pFalse
}
