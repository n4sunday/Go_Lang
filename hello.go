package main

import ("fmt")

//Structure

type Books struct{
	title string
	author string
	subtitle string
	price float64	
}
func main() {
	/*var Book1 Books
	Book1.title="Go Programming"
	Book1.author="Nattapon"
	Book1.subtitle="N4S"
	Book1.price=199
	fmt.Println(Book1)
	fmt.Println(Book1.title)*/

	Golang:=Books{title:"Go Programming",author:"Nattapon",price:199}
	fmt.Println(Golang)
	fmt.Println((Golang.price*20)/100)


}