package main
import (
	"fmt" 

)

//While Go doesn’t have inheritance, it encourages code reuse via built-in support for composi‐tion and promotion

type Employee struct {
	Name string
	ID string
   }
   func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
   }
   type Manager struct {
	Employee					//Her no name is assigned to the Employee so we call it embedded field
	Reports []Employee
   }
//    func (m Manager) FindNewEmployees() []Employee {
// 	// do business logic
//    } 

//Any field or methods declared on an embedded field are promoted to the containing struct and can be invoked directly on it 



type Inner struct {
	X int
   }
   type Outer struct {
	Inner
	X int
   }

func main11(){

	o := Outer{
		Inner: Inner{
		X: 10,
		},
		X: 20,
	   }
	   fmt.Println(o.X) // prints 20
	   fmt.Println(o.Inner.X) // prints 10

}