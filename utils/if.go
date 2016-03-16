package utils

func If(condition bool, pTrue interface{}, pFalse interface{}) interface{} {
	if (condition) {
		return pTrue;
	}
	return pFalse;
}