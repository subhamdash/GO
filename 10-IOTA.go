package main
import (
	"fmt" 

)

/*
Many programming languages have the concept of enumerations, where you can
specify that a type can only have a limited set of values. Go doesn’t have an enumera‐
tion type. Instead, it has iota, which lets you assign an increasing value to a set of
constants.


iota is commonly used in Go for defining constants that form a sequence or a pattern, such as days of the week, bit flags, 
or enum-like structures. It simplifies the code and makes it easier to maintain when you need consecutive values for a series of constants.

*/

const (
    A = iota // A = 0
    _        // Skip (value discarded)
    B        // B = 2 (implicit assignment of iota)
    C        // C = 3 (implicit assignment of iota)
)

func main10(){
	fmt.Println(A, B, C)

}