package main

import (
	"flag"
	"fmt"
)

//The flag package in Go is used to develop a UNIX system-like program that accepts
//command-line arguments to manipulate its behavior in some form

// THink of args and argv
func main20() {

	// str1 := flag.String("u", "root", "username")
	// str2 := flag.String("pass", "", "password")
	// flag.Parse()
	// fmt.Println("Username : ", *str1)
	// fmt.Println("Password : ", *str2)

	var hostname, port string
	flag.StringVar(&hostname, "h", "127.0.0.1", "hostname")  //It returns a pointer to a string variable.
	flag.StringVar(&port, "p", "8081", "port")
	flag.Parse()
	fmt.Println("The Port is ", port)

}

//go run "/Users/subhamdash/Desktop/Go/LoadBalancer/be/BakEnd.go" -u sdash32 -p 1234
