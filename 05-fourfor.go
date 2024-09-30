package main
import (
	"fmt" 

)
func main05(){

	//A complete

	// First, you must use := to initialize the variables; var is not legal here. 
	//Second,just like variable declarations in if statements, you can shadow a variable here.
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	   }



	//• A condition-only for
	x := 1
	for x < 100 {
	fmt.Println(x)
	x = x * 2
	}

	//A infininte For

	// for {
	// 	fmt.Println("Hello")
	// 	}

	//For-Range Statement

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
	fmt.Println(i, v)
	}

}