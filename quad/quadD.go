package main

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	for i := 1; i <= y; i++ { //(i pour les colonnes)
		for j := 1; j <= x; j++ { //(j pour les lignes)
			if (i == y && j == 1) || (i == 1 && j == 1) { //(1ere case || derniere case -- 1ere colonne)
				z01.PrintRune('A')
			} else if (i == y && j == x) || (i == 1 && j == x) { //(1ere case || -- derniere colonne)
				z01.PrintRune('C')
			} else if ((i == 1 || i == y) && !(j == 1 || j == x)) || ((j == 1 || j == x) && !(i == 1 || i == y)) {
				z01.PrintRune('B') //(les 1ere et derniere ligne sauf les cases des extremités || les 1ere et derniere colonnes sauf les cases des extremités  )
			} else {
				z01.PrintRune(' ') // pour les espaces au milieu du rectangle
			}
		}
		z01.PrintRune('\n')
	}
}
