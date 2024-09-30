package main
import (
	"fmt" 
	"math/rand"
)
func main04(){
	//we can see the variables are scoped to the conditon that is for both if and else

	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	   } else if n > 5 {
		fmt.Println("That's too big:", n)
	   } else {
		fmt.Println("That's a good number:", n)
	   }
}
