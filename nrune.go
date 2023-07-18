package piscine

func NRune(s string, n int) rune {
	count := 1
	for _, r := range s {
		if count == n {
			return r
		}
		count++
	}
	return 0
}
