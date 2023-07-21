package piscine

func SplitWhiteSpaces(s string) []string {
	var w []string
	a := 0
	for i, r := range s {
		if r == ' ' || r == '\n' || r == '\t' {
			w = append(w, s[a:i])
			a = i + 1
		}
	}
	w = append(w, s[a:])
	return w
}
