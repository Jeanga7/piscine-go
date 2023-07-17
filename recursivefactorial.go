package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	r := 1
	if nb > 1 {
		r = nb * RecursiveFactorial(nb-1)
		if r < 0 { // depassement
			return 0
		}
	}
	return r
}
