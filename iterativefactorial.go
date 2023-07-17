package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 { // exclure les nbrs negatif (valeurs impossibles)
		return 0
	}
	r := 1
	for i := 1; i <= nb; i++ {
		r *= i     // pour avoir le factoriel(1*2*...*nb)
		if r < 0 { // depassement
			return 0
		}
	}
	return r
}
