package main
import (
    "fmt"

)


func main25() {
    ch := make(chan int) // Unbuffered channel

    // Write to the channel in the main goroutine
    ch <- 10 // This line will block until there's another goroutine ready to receive from the channel

    // Read from the channel in the main goroutine
    value := <-ch // This line will block until there's another goroutine ready to send data into the channel

    fmt.Println("Received value:", value)
}
/*
The main function is also a goroutine.
ch <- 10 is writing to a channel and is done by main goroutine
it will will block until another go routine ready to receive the data from the channel.
Insort we need atleat tow concurrently go routine required to play around the channel.
Read and write are too vauge to descripe the better manner would be send and recieve .


communication via channels enforces synchronization between goroutines, 
ensuring safe data exchange and coordination between concurrent parts of the program.

*/
/*
By default channels are unbuffered. Every write to an open, unbuffered channel
causes the writing goroutine to pause until another goroutine reads from the same
channel


Likewise, a read from an open, unbuffered channel causes the reading
goroutine to pause until another goroutine writes to the same channel. This means
you cannot write to or read from an unbuffered channel without at least two concur‐
rently running goroutines.

*/
