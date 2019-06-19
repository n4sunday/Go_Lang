package main


import (
	"fmt"
)

var pow = []int{1, 2,4,8,16,32,64,128}
var name = []string{"Sun","Rob","John"}

func main(){
	for i, v:=range pow {
		fmt.Printf("2**%d = %d\n",i,v)
	}

	for i, x:=range name {
		fmt.Printf("%d %s \n",i,x)
	}
}