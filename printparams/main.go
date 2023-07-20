package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	n := os.Args[1:]
	for _, i := range n {
		for _, s := range i {
			z01.PrintRune(s)
		}
		z01.PrintRune('\n')
	}
}
