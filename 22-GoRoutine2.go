package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(100 * time.Millisecond) // Simulate some work
    }
}

func main22() {
    go printNumbers() // Start printNumbers in a new goroutine

    // Continue executing the main program
    for i := 1; i <= 5; i++ {
        fmt.Println("Main:", i)
        time.Sleep(200 * time.Millisecond) // Simulate some work
    }

    // Wait for a while to allow the goroutine to finish
    time.Sleep(1 * time.Second)
}

//But can we get some return value in the go routine function , No right so we need to have a channel to commumicate 
