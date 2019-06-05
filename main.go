package main
import (
	"fmt"
)

func Total(x, y string) (string,string) {
	return y,x
}

func main(){
	a,b := Total("world","Hello")
	fmt.Println(a,b)
}