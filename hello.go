package main

import (
	"fmt"
)

//Input Scanf
func main() {
	fmt.Println("Input You Number:")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 10
	fmt.Println("Sum:", output)

}
