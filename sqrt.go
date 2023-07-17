package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	for i := 0; i <= nb; i++ {
		if nb == i*i {
			return i
		}
	}
	return 0
}
