package main

import (
	"fmt"
)

//Scope

//Global Varaible
var gVaraible int = 500

func main() {
	lVaraible := 40
	fmt.Println("Global", gVaraible)
	fmt.Println("Local", lVaraible)
	anotherFunction()
}
func anotherFunction() {
	fmt.Println("Global", gVaraible)
	//fmt.Println("Local", lVaraible)
}
