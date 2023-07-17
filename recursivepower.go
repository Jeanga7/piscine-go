package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	r := 1
	for i := 1; i <= power; i++ {
		r = nb * RecursivePower(nb, power-1)
	}
	return r
}
