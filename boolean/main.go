package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	lengthOfArg := 0
	for i := range os.Args[1:] {
		lengthOfArg = i + 1
	}
	if isEven(lengthOfArg) {
		PrintStr(EvenMsg)
	} else {
		PrintStr(OddMsg)
	}
}
