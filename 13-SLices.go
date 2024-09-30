
package main
 
import "fmt"
 
func main13() {
    array := [5]int{1, 2, 3, 4, 5}
    slice := array[1:4]
	slice1 := []int{1, 2, 3}
    slice1 = append(slice1, 4, 5, 6)
 
    fmt.Println("Array: ", array)
    fmt.Println("Slice: ", slice)
	fmt.Println("Slice1: ", slice1)
	fmt.Println("Slice1: ", slice1)

	var my_slice[]int //Declaration of slice
	my_slice = append(my_slice, 4, 5, 6)
	fmt.Println("my_slice: ", my_slice)

	var my_slice_1 = []string{"Geeks", "for", "Geeks"} //Declare and Intialize the slice
	fmt.Println("my_slice1: ", my_slice_1)

}

/*Slice are variable lenght sequence 


Pointer: The pointer is used to points to the first element of the array that is accessible through the slice. 
Here, it is not necessary that the pointed element is the first element of the array.
Length: The length is the total number of elements present in the array.
Capacity: The capacity represents the maximum size upto which it can expand.
*/