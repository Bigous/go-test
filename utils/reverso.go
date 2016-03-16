package strings

// Reverso retorna uma string com os caracteres invertido em relação a string passada.
func Reverso(s string) string {
	var (
		r []rune;
		i, j, t int;
	);
	r = []rune(s);
	t = len(r);
	for i, j = 0, t-1; i < t/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i];
	}
	return string(r);
}
