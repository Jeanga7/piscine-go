package piscine

func NRune(s string, n int) rune {
	for pos, r := range s {
		if pos+1 == n {
			return r
		}
	}
	return 0
}
