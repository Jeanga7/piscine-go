package piscine

func ToUpper(s string) string {
	j := []rune(s)
	for i, up := range j {
		if up >= 'a' && up <= 'z' {
			j[i] = up - ('a' - 'A')
		}
	}
	return string(j)
}
