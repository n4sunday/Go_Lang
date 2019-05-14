package main

import (
	"fmt"
)

//Loop
func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Sunday #",i)
	}

	j:=1
	for j<=20 {
		if(j%2==0){
			fmt.Println("Even",j)
		}else {
			fmt.Println("Odd",j)
		}
		j++
	}
}
