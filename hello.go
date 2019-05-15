package main

import ("fmt")

//Closure (Function No Name)
func main() {
	add:=func (x,y int) int{
		return x+y
	}
	fmt.Println(add(5,10))
	x:=0
	increment:=func() int{
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())


}