package main

import (
	"fmt"
	"os"
)

func main() {
	alert := []string{"01", "galaxy", "galaxy 01"}
	arg := os.Args[1:]
	for _, i := range alert {
		for _, j := range arg {
			if i == j {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
