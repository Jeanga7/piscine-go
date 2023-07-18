package main

import (
	"fmt"
)

func Index(s string, toFind string) int {
	t := len(toFind)
	for i := range s {
		if i+t <= len(s) && s[i:i+t] == toFind {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
