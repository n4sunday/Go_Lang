package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z string
	A float64
	B bool
}
type Num struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1,2,"Sun",3.14,true})

	v := Num{10, 20}
	fmt.Println(v.X)
	fmt.Println(v.Y)
}