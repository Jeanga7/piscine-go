package piscine

func RecursiveFactorial(nb int) int {
	r := 1
	if nb < 0 || nb > 99 {
		return 0
	}
	if nb > 1 {
		r = nb * RecursiveFactorial(nb-1)
	}
	return r
}
