package piscine

func LastRune(s string) rune {
	var Lastword rune
	for _, r := range s {
		Lastword = r
	}
	return Lastword
}
