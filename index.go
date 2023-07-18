package main

import (
	"fmt"
	"strings"
)

func Index(s string, toFind string) int {
	return strings.Index(s, toFind)
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
