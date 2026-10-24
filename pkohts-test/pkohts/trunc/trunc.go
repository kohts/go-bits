package main

import "fmt"

func main() {
	var aFloat float32
	fmt.Printf("Please enter a floating point number: ")
	fmt.Scan(&aFloat)
	fmt.Printf("Got: %d\n", int(aFloat))
}
