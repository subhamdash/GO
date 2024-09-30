package main
import (
	"fmt" 


)

//Go doesn't have classes, because it doesn't have inheritance 
//Decalring a structure
type Person struct{
	name string
	age int
	pet string
}


func main07(){
	var fred Person
	bob:=Person{}
	julia:= Person{
		"Julia",
		26,
		"cat",
	}
	beth:= Person{
		age:30,
		name:"Beth",
	}

	//Anonymus Struct
	//The first is when you translate external data into a struct or a struct into exter‐nal data (like JSON or protocol buffers). 
	//This is called unmarshaling and marshaling data.
	pet := struct {
		name string
		kind string
	   }{
		name: "Fido",
		kind: "dog",
	   }

	fmt.Println(fred,bob,julia,beth,pet)


	//WE cannot compare the strucuts or primitive data types in the GO but we can do it for the anonymus struct.

}