package piscine

func TrimAtoi(s string) int {
	r := 0 // pour stocker le resultat final
	x := 1 // pour le signe (1 pour +; -1 pour -)
	// parcours de la chaine s
	for _, c := range s {
		if c >= '0' && c <= '9' {
			r = r*10 + int(c-'0')
		}
		if c == '-' {
			x *= -1
		}
	}
	return r * x
}
