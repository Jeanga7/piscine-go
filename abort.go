package piscine

func Abort(a, b, c, d, e int) int {
	tab := []int{a, b, c, d, e}
	for i := 0; i <= 4; i++ {
		for j := i + 1; j < 5; j++ {
			if tab[j] < tab[i] {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}
	}
	return tab[2]
}
