package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	des := true
	asc := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			asc = false
		}
		if f(a[i], a[i+1]) > 0 {
			des = false
		}
	}
	return des || asc
}
