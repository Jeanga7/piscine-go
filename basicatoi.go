package piscine

func BasicAtoi(s string) int {
	basic := 0
	for _, atoi := range s {
		basic = basic*10 + int(atoi-'0')
	}
	return basic
}
