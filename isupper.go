package main

import "fmt"

func IsUpper(s string) bool {
	for _, i := range s {
		if !(i >= 'A' && i <= 'Z') {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))

}
