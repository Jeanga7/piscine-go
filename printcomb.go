package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for k := '1'; k <= '8'; k++ {
			for q := '2'; q <= '9'; q++ {
				if i < k {
					if k < q {
						z01.PrintRune(i)
						z01.PrintRune(k)
						z01.PrintRune(q)
						if i < 7 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						} else {
							z01.PrintRune('\n')
						}
					}
				}
			}
		}
	}
}
