package piscine

func ToLower(s string) string {
	low := []rune(s)
	for i, j := range low {
		if j >= 'A' && j <= 'Z' {
			low[i] = j - ('A' - 'a')
		}
	}
	return string(low)
}
