package main

/*
he bufio package in Go provides two main types for buffered I/O operations: bufio.Reader and bufio.Scanner. 
While both are useful for reading data from an io.Reader, they serve slightly different purposes and have different behaviors:

bufio.Reader:

bufio.Reader is a type that wraps an io.Reader and provides buffered reading operations.
It's designed for reading raw data efficiently, particularly when you know the format of the data and want to read it byte by byte, 
line by line, or using custom delimiters.
It doesn't provide any built-in parsing capabilities; instead, it offers low-level methods 
like Read, ReadByte, ReadLine, and ReadSlice for reading data in various formats.
bufio.Scanner:

bufio.Scanner is a type that wraps an io.Reader and provides a convenient way to tokenize data.
It's designed for parsing structured textual data, such as text files with lines or fields separated by delimiters.
It automatically splits input into tokens (words, lines, etc.) based on a split function, 
which is typically a custom function specified by the user or one of the predefined split functions like bufio.ScanLines or bufio.ScanWords.
It's useful for tasks like reading and parsing lines from log files, processing CSV files, or extracting data from structured text.
Here's a brief comparison:

Use bufio.Reader when you need low-level control over reading raw data efficiently, byte by byte or using custom logic.
Use bufio.Scanner when you want a higher-level interface for tokenizing and parsing structured textual data, 
such as lines or fields separated by delimiters.
In summary, bufio.Reader is more versatile and suitable for general-purpose buffered reading, 
while bufio.Scanner is more specialized and suited for parsing structured text. 

Choose the appropriate one based on your specific requirements and the nature of the data you're working with.








*/


import(
	"fmt"
	"flag"
	"os"
	"bufio"
)

func main36(){

	dict := flag.String("dict", "/Users/subhamdash/Desktop/Go/BloomFilter/dict.txt", "dictionary file")
	// It returns a pointer to a string variable.

	file_content,err := os.Open(*dict)
	if err!=nil {
		fmt.Println("There is an error",err)
	} else{
	

		//scanner := bufio.NewScanner(file_content)
		reader:=bufio.NewReader(file_content)

		// for scanner.Scan() {
		// 	line := scanner.Text()
		// 	fmt.Println(line)
		// }
	

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break // Reached end of file
			}
			fmt.Print(line)
		}

		

	}

}