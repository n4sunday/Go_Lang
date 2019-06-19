package main

import ( 
	"fmt"
)

func main(){
	var s []int 
	printSlice(s)

	//appendworks on nil slices.
	s = append(s, 0)
	printSlice(s)

	//The slice grows as needed.
	s = append(s,1)
	printSlice(s)

	//We can add more then noe element at atime.
	s = append(s, 2,3,4)
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n",len(s), cap(s), s)
}