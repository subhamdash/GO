package main
import (
	"fmt" 
	"errors"

)

/*
error is a built-in interface that defines a single method:
type error interface {
 Error() string
}

Anything that implements this interface is considered an error. The reason why we
return nil from a function to indicate that no error occurred is that nil is the zero
value for any interface type*/

//First way to handel the error
func doubleEven(i int) (int, error) {
	if i % 2 != 0 {
	return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

//Second way is to call fmt.Errorf function

func doubleEven2(i int) (int, error) {
	if i % 2 != 0 {
	return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
   }


