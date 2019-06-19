package main


import "fmt"

func main(){
	var s []int
	fmt.Println(s, len(s), cap(s))
	fmt.Printf("%d len:%d cap:%d \n",s, len(s), cap(s))
	if s == nil {
		fmt.Println("S is nil!")
	}
}