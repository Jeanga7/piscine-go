package piscine

func SplitWhiteSpaces(s string) []string {
	var w []string
	a := ""
	for _, r := range s {
		if string(r) != " " {
			a += string(r)
		}
		if r == ' ' || r == '\n' || r == '\t' {
			if a != "" {
				w = append(w, string(a))
				a = ""
			}
		}
	}
	w = append(w, a)
	return w
}
