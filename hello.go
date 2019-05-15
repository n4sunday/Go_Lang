package main

import ("fmt")

//Go Routine Function

func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)
}
func f(n int){
	for i:=0; i<100; i++{
		fmt.Println(n,":",i)
	}
}