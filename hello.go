package main

import ("fmt")

//Variadic Fucntion
func main() {
	summation(1,2)
	summation2(1,2,3)
	sumVariadic(1,2,3,4,5,6,7,8,9,10)
}

//Basic
func summation(a int,b int){
	fmt.Println(a+b)
}
func summation2(a int,b int,c int){
	fmt.Println(a+b+c)
}

//Variadic 
func sumVariadic(nums...int){
	var total int
	for _ , n:=range nums{
		total+=n
	}
	fmt.Println(total)
}