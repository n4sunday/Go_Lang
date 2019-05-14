package main

import ("fmt")

//Recursive Fucntion
func main() {
	fmt.Println(factorial(5))
}

func factorial(num int) int{
	if num==0{
		return 1
	}
	return num*factorial(num-1)

}