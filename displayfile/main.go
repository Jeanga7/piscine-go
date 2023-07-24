package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// recuperer les argument de laligne de commande
	arg := os.Args[1:]
	// verifier le nombre d'argumment
	if len(arg) == 0 {
		fmt.Println("File name missing")
	} else if len(arg) > 1 {
		fmt.Println("Too many arguments")
	} else {
		// lire le contenue du fichier
		if arg[0] == "quest8.txt" {
			contenue, erreur := ioutil.ReadFile(arg[0])
			if erreur != nil {
				fmt.Println(erreur)
				return
			}
			fmt.Println(string(contenue))
		}

	}
}
