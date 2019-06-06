package main

import (
	"fmt"
)

type Num struct {
	X int
	Y int
}

func main(){
	v := Num{10,20}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}