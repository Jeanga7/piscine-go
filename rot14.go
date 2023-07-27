package piscine

func Rot14(s string) string {
	r := []rune(s)
	for i, val := range s {
		if (val >= 'a' && val <= 'z') || (val >= 'A' && val <= 'Z') {
			if val >= 'm' && val <= 'z' || val >= 'M' && val <= 'Z' {
				r[i] = val - 12
			} else {
				r[i] = val + 14
			}
		}
	}
	return string(r)
}
