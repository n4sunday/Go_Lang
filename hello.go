package main

import ("fmt")

//Array
func main() {
	var x[5] int
	x[0]=5
	x[1]=10
	x[2]=15
	x[3]=20
	x[4]=25
	fmt.Println(x)
	fmt.Println(x[1])

	y:=[5]float64 {5,4,6,2,1,}
	fmt.Println(y)

	var total float64
	for _ , value:=range y {
		total += value
	}
	fmt.Println("total",total)
	fmt.Println("AVG total",total/float64(len(y)))

}