package main

import ("fmt")

//Fucntion
func main() {
	dosome()
	dosomething("Sunday")
	addition(8,2)
	result:=additio2(15,15)
	fmt.Println("result",result*10)
}

func dosome(){
	fmt.Println("Hello Word")
}

func dosomething(str string){
	fmt.Println("Func : ",str)
}
func addition(a int ,b int){
	fmt.Println(a+b)
}
//Return 
func additio2(a int ,b int) int{
	output:=a+b
	return output
}