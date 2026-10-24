package main

import "fmt"

func changeSlice(aSlice []int) {
	aSlice[0] = 3
}

func main() {
	aSlice := []int{1, 2, 3}
	fmt.Printf("%v", aSlice)
	changeSlice(aSlice)
	fmt.Printf("%v", aSlice)
}
