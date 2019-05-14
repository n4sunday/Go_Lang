package main

import (
	"fmt"
)

func main() {
	//Operater Math
	fNumber := 60
	fSecond := 5
	fmt.Println(fNumber + fSecond)
	fmt.Println(fNumber - fSecond)
	fmt.Println(fNumber * fSecond)
	fmt.Println(fNumber / fSecond)
	//Strung
	p1 := "Nattapon"
	p2 := "Sunday"
	//Concatenation
	p3 := p1 + p2
	fmt.Println(p3)
	fmt.Println(p3[0:3]) //Nat
	fmt.Println(p3[0:])  //NattaponSunday
	fmt.Println(p3[0])   //Ascii
}
