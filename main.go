package main 
import ( "fmt")

func main(){
	sum := 1
	for sum<20 {
		sum += sum
		fmt.Println(sum)
	}	
}
//while in Go delete ; ;