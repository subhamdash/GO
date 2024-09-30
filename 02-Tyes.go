
package main
import "fmt"
func main02() {
	var x int;
	var y int=10;
	var z=10;
	
	//All these above are same 
	var a,b=10,"Hellow";
	
	c:=100 //:= replaces declaration of var
	fmt.Println(x,y,z,a,b,c)
	//We cannot leave var x as we not specifying anything whether it's a int string or what

	var g [3] int;
	var gh=[3]int{10,20,30}
	var xy = [12]int{1, 5: 4, 6, 10: 100, 15} //we are creating a sparse array
	fmt.Println(g,gh,xy)


	//var nilMap map[string]int  Maps in Go
	//totalWins := map[string]int{}



}

