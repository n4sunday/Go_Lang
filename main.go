package main
import (
	"fmt"
	"time"
)
func main(){
	t := time.Now()
	fmt.Printf("Time now : %s\n",t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("God afternoon.")
	default:
		fmt.Println("Good evning.")
	}
}