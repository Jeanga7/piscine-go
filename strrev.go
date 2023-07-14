package piscine

func StrRev(s string) string {
	rev := ""
	for _, i := range s {
		rev = string(i) + rev
	}
	return rev
}
