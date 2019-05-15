package main

import ("fmt")

//Pointer
func main() {
	x:=10
	fmt.Println("x value is ",x)
	fmt.Printf("x value is %d \n",x)
	fmt.Printf("Address x variable %x\n",&x)

	var p *int
	p=&x //Pointer at x
	fmt.Printf("Pointer P %x \n",p)
	*p = 20
	fmt.Println("p value is ",p)
	fmt.Println("x value is ",x)



}