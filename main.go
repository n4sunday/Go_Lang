package main

import (
	"fmt"
)

type Num struct {
	X int
	Y int
}

var (
	v1 = Num{1 ,2} //has type Num
	v2 = Num{X: 10} //Y:0 is imlicit
	v3 = Num{}     //X:0 and Y:0
	p = &Num{1,2}  //has type *Num
)

func main() {
	fmt.Println(v1,p,v2,v3) //{1 2} &{1 2} {10 0} {0 0}
}


