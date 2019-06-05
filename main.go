package main 
import ( 
	"fmt"
	"math"
)

func pow(x, n, a float64)(num float64){
	if v := math.Pow(x,n); v < a{
		return v
	}
	return a
}

func main(){
	fmt.Println(pow(3,2,10))
	fmt.Println(pow(3,3,20))
}