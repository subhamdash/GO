package main
import (
	"fmt"
	"time"

)
func attack(target string){
	fmt.Println("Throwing stars at ",target)
	time.Sleep(time.Second)
}

func main21(){
	start:= time.Now()
	defer func(){
		fmt.Println(time.Since(start))
	}()

	evilNinjas:= []string{"Tommy","Jhony","Bobby","Andy"}
	for _,evilNinja := range evilNinjas{
		go attack(evilNinja)   //WE can just have go keyword but we are not getting the correct result because it is exiting out 
	}
	fmt.Println("All stare are Thrown")
	time.Sleep(time.Second*2)

}

//Here the output will vary because the Execution of different thread are occuring concurrently so there can be possibility that
//Tommy may be first of Last to return similarly with all the cases with all the 24 cominantions possible 

//Placing the go keyword before a function call in Go creates a new goroutine to execute that function concurrently with the rest of the program. It initiates a separate, independent flow of execution, allowing the function to run concurrently with the main program or other goroutines.
/*
Here's how it works:

Goroutines: Goroutines are lightweight threads managed by the Go runtime. 
They enable concurrent execution of code within a single process. 
When you use go before a function call, Go creates a new goroutine to execute that function, allowing it to run concurrently with the rest 
of the program.

Non-blocking: Using go doesn't block the execution of the current goroutine. 
The main program or the current goroutine continues execution immediately after spawning the new goroutine, 
without waiting for the spawned goroutine to finish.

Independence: The new goroutine executes independently of the main program or other goroutines. 
It has its own call stack and executes concurrently with other goroutines, potentially on different CPU cores.

*/