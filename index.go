package piscine

func Index(s string, toFind string) int {
	t := len(toFind)
	for i := range s {
		if i+t <= len(s) && s[i:i+t] == toFind {
			return i
		}
	}
	return -1
}
