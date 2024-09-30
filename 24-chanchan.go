package main

import (
    "fmt"

)
//a := <-ch // reads a value from ch and assigns it to a
//ch <- b // write the value in b to ch


func produceData(ch chan<- int) {   //Write data to the channel 
    for i := 1; i <= 5; i++ {
        ch <- i // Send data on the channel
    }
    close(ch) // Close the channel when done
}


func consumeData(ch <-chan int) {  //Read data from the channel
    for num := range ch {
        fmt.Println("Received:", num)
    }
}

func main24() {

    fmt.Println("Hellow world")
	ch:=make(chan int) //Make the channel , but it's unbuffered one 
	/*
	ch <- 10
	a:= <- ch
	fmt.Println(a)
	close(ch)
	*/
	//WE cannot write th program like this as the channels are unbuffered, 
	//We have two ways to do a thing either  use a goroutine or use a buffereed channel
	go produceData(ch) //So we are using a go routine to produce the ch and we will get the data from it.
	//consumeData(ch)
	



	


}
/*
By default channels are unbuffered. Every write to an open, unbuffered channel
causes the writing goroutine to pause until another goroutine reads from the same
channel


Likewise, a read from an open, unbuffered channel causes the reading
goroutine to pause until another goroutine writes to the same channel. This means
you cannot write to or read from an unbuffered channel without at least two concur‐
rently running goroutines.

*/



/* Channels in Go are communication mechanisms that allow goroutines to synchronize execution and communicate by sending and receiving values.
Channels facilitate safe communication and coordination between concurrent goroutines, 
enabling them to exchange data and signals without race conditions or explicit locking.

In the context of using channels with goroutines created by the go keyword, channels play a crucial role in coordinating the execution 
of concurrent goroutines and facilitating communication between them. Here's how channels help:

Synchronization: Channels can be used to synchronize the execution of concurrent goroutines. By sending or receiving on a channel, 
				 goroutines can coordinate their actions to ensure certain operations happen in a specific order.

Data Exchange: Channels allow goroutines to exchange data safely. One goroutine can send data on a channel, 
			   and another goroutine can receive that data, enabling communication and data sharing between goroutines.

Blocking Operations: Sending or receiving data on a channel is a blocking operation. If a goroutine attempts to send data on a channel and 
					 there's no goroutine ready to receive it, the sender will block until another goroutine is ready to receive. 
					 Similarly, if a goroutine attempts to receive data from an empty channel, it will block until another goroutine 
					 sends data.


*/
