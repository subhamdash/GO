package main
import (
    "fmt"

)


func main26() {
    ch := make(chan int,10) // Unbuffered channel, here are defining the capacity , Here we say that here we are writing without blocking.


	for i:=0;i<10;i++ {  
	ch <- i;
	}
	close(ch)  //WE need to close the channel so to make sure other will not wait for them to write

	

    for v:= range ch{
   // This line will block until there's another goroutine ready to send data into the channel

    fmt.Println("Received value:", v)
	}
	
}

/*
The program you provided will exit because it falls into a deadlock situation. 
Although you've created a buffered channel with a capacity of 10, which means it can hold up to 10 values without blocking, 
the issue arises when you attempt to range over the channel without any additional goroutines to send more values.

Let's break down what's happening:

1. You create a buffered channel ch with a capacity of 10.
2. You then loop from 0 to 9 and send each value into the channel ch. This works fine because the channel has sufficient buffer space 
   to hold these values without blocking.
3. After sending all the values, you attempt to range over the channel ch. 
   This loop will iterate over the buffered values in the channel until the channel is closed. 
   However, since no additional goroutines are sending more values into the channel, 
   the range loop will eventually exhaust all buffered values and then block waiting for more values to be sent.


Since there are no additional goroutines sending values into the channel, and the range loop is blocked waiting for more values, 
the program enters a deadlock state, and it will not exit gracefully.

To avoid deadlock, you need to ensure that the range loop has an exit condition or that there are goroutines 
sending values into the channel. For example, you could close the channel after sending all values or use a separate goroutine to 
send values while the main goroutine ranges over the channel.
*/