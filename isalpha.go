package piscine

func IsAlpha(s string) bool {
	for _, i := range s {
		if !(i >= 'a' && i <= 'z' || i >= 'c' && i <= 'Z' || i >= '0' && i <= '9' || i == ' ') {
			return false
		}
	}
	return true
}
