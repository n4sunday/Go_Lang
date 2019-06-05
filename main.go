package main
import (
	"fmt"
)

func spit(num int)(x int, y int){
	x = num * 5 / 2
	y = num - 10
	return 
}

func main() {
	fmt.Println(spit(20))
}

