package main
import (
    "fmt"
	//"time"

)
func check_write(ch chan<-int){
	ch <- 20
}

func read(ch2 <-chan int){
	a:= <-ch2
	fmt.Println(a)


}
/*THis code won't give us a deadlock as one goroutine will block  till another go routine is ready to recieve the data

func main() {
	ch2 := make(chan int)
	go func(){
		a:= <-ch2
		fmt.Println(a)

	}()
	ch2<-10
}
*/

/* this will also will not cause any error as there is two go routine and it just sends the value it doesn't need to wait for reciever

func main() {
	ch2 := make(chan int)
	go func(){
		ch2<-11
	

	}()
	
}*/

/* here there is no two goroutines so there will be a deadlock for sure as it will wait to send the data but there is no reciever
func main() {
	ch2 := make(chan int)
	go func(){
		
	

	}()
	ch2<-11
}
*/




func main27() {
    ch := make(chan int) 
	ch2 := make(chan int)

	go func(){
		ch <- 10; 
		a:= <-ch2
		fmt.Println(a)

	}()
	ch2<-10
	a:=<-ch
	fmt.Println(a)


}


/* A dead lock occurs becuase main go routine waits for ch and the clousre go function wait for ch2 , so both go routine wait to recieve
and send data and this infinte wait will give us deadlock
This code snippet has the potential to result in a deadlock. Let's break down the execution flow:

The main goroutine creates two unbuffered channels, ch and ch2.
It launches a new goroutine using an anonymous function to perform the following:
Send the value 10 into channel ch.
Receive a value from channel ch2 and print it.
In the main goroutine, it:
Sends the value 10 into channel ch2.
Receives a value from channel ch and prints it.
Potential Deadlock Scenario:

The main goroutine sends 10 into ch2.
The anonymous goroutine sends 10 into ch.
However, the anonymous goroutine is blocked on receiving from ch2 (a := <-ch2) because no value has been sent into ch2 yet.
Similarly, the main goroutine is blocked on receiving from ch (a := <-ch) because no value has been sent into ch yet.
Both goroutines are waiting for each other to proceed, leading to a deadlock.

*/



	// select
	// {
	// case ch2<-10: //This main go routine will block until another go routine recieves the data
	// //b:= <- ch;
	// }
	
	
	//fmt.Println(b)