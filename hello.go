package main

import (
	"fmt"
)

//If Else Statment
func main() {
	fmt.Println("Input You Number:")
	var input float64
	fmt.Scanf("%f", &input)
	condition := input > 2
	if condition {
		fmt.Println("Worked")
	} else {
		fmt.Println("Not Worked")
	}

	if 6 > 3 && 8 > 5 {
		fmt.Println("Worked 6>3 && 8>5")
	} else {
		fmt.Println("Not Worked")
	}

	if 6 > 3 || 3 > 5 {
		fmt.Println("Worked 6>3 || 8>5")
	} else {
	}

}
