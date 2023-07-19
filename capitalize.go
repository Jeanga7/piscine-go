package main

import "fmt"

func Capitalize(s string) string {
	rs := []rune(s) // chaine en rune
	sc := true      // pour les maj
	for i, r := range rs {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			if sc && r >= 'a' && r <= 'z' { //
				rs[i] = r - 'a' + 'A' // pour mettre en maj
			}
			if !sc && r >= 'A' && r <= 'Z' {
				r[i] = r - 'A' + 'a' // pour les enlever les maj qui devez pas etre maj
			}
			sc = false
		} else {
			sc = true
		}
	}
	return string(rs) // convers rune - string et renvoyer
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
