package main
import (
	"fmt" 
	"errors"

)
//The things we need to carefully observer is that the params declared and return type is placed after 
//closing paranthesis and opening braces
//If your function return something then we need to have a return type or else not needed.


func div(numerator int, denominator int) int {
	if denominator == 0 {
	return 0
	}
	return numerator / denominator
   }

//We are returnin multiple values in go . We need to seperate them by commas not put a parathesis around them.
//You must assign each value returned from a function. If you try to assign multiple return values to one variable, you get a compile-time error.
func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
	return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
   }

func main06(){
	res:= div(10,1)
	fmt.Println(res)

	/*
	Go supports variadic parameters. The variadic parameter must be
the last (or only) parameter in the input parameter list. You indicate it with three dots
(…) before the type. The variable that’s created within the function is a slice of the
specified type */



//Go is call by value , hence the changes made inside the function won't be effective outside the function

//This is different for Maps and slices due to their implementaitoj using pointer. Every type in Go is a vlaue type. 
//It's just that sometimes the value is a pointer.

type opFuncType func(int,int) int

//Anonymus Function to be used
		for i := 0; i < 5; i++ {
			func(j int) {
				fmt.Println("printing", j, "from inside of an anonymous function")
			}(i)
		}
}