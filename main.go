package main

import ("fmt")

func main(){
	i, j := 5, 10

	p := &i				 //point to i
	fmt.Println(*p) //5  //read i through the poiner
	*p = 4				 //set i through the pointer
	fmt.Println(i)	//4	 //see the new value of i

	p = &j				 //point to j
	*p = *p/5 			 //divide j through the pointer
	fmt.Println(j)  //2	 //see the new value of j
}