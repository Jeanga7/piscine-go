package piscine

func BasicAtoi2(s string) int {
	basic := 0
	for _, i := range s {
		if i < '0' || i > '9' {
			return 0
		}

		basic = basic*10 + int(i-'0')
	}
	return basic
}
