package main
import "fmt"
func main03(){
	x := 10
 if x > 5 {
 fmt.Println(x)
 x := 5
 fmt.Println(x)
 }
 fmt.Println(x)

 /*A shadowing variable is a variable that has the same name as a variable in a contain‐
ing block. For as long as the shadowing variable exists, you cannot access a shadowed
variable.*/


/*
a := 10
 fmt.Println(a)
 fmt := "oops"
 if a>=10{
 fmt.Println(fmt)
 }
*/  // THis is giving error as we are trying to shadow the fmt package which makes it impossible to make fmt package for the rest of the main
//function
//We can avoid it by using linters in out code


}