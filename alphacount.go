package piscine

func AlphaCount(s string) int {
	count := 0
	for _, val := range s {
		if val >= 'a' && val <= 'z' || val >= 'A' && val <= 'Z' {
			count++
		}
	}
	return count
}
