package main

import (
	
	"github.com/01-edu/z01"
)
func main() {
	PrintNbr(-9223372036854775808)
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}


func PrintNbr(n int) {

	t:=1
		if n<0{
			t=-1
			z01.PrintRune('-')
		}
		if n!=0{
			f:=(n/10)*t
			if f!=0{
				PrintNbr(f)
			}
			k:=((n%10*t))+'0'
			z01.PrintRune(rune(k))
			}else{
				z01.PrintRune('0')
			}
		
}