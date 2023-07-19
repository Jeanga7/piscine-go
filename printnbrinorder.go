package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	s := [10]int{} // tableau de taille 10 pour compter de 0 a 9
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for n > 0 {
		d := n % 10 // obtenir le dernier chiffre de n
		s[d]++      // incrementer le compteur pour ce chiffre
		n /= 10     // suppr dernier chiffre de n
	}
	for i := 0; i < 10; i++ {
		count := s[i] // recuperer le compteur pour i
		for count > 0 {
			z01.PrintRune(rune(i + '0')) // imprimer le chiffre i en tant que caractere
			count--                      // decrementer le compteur
		}
	}
}
