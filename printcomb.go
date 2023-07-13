package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for k := '1'; k <= '8'; k++ {
			for q := '2'; q <= '9'; q++ {
				if i < k && k < q {
					z01.PrintRune(i)
					z01.PrintRune(k)
					z01.PrintRune(q)
				}
			}
		}
	}
	z01.PrintRune('\n')
}
