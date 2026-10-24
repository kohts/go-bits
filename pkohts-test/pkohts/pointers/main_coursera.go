package main

import "fmt"

func main() {
	v1 := "regular variable"
	v2 := &v1
	v3 := &v1

	fmt.Println("Hello, world!")

	fmt.Println("regular variable value:", v1)
	fmt.Println("pointer to regular variable:", *v2)

	if v1 == *v2 {
		fmt.Println("regular variable equals to its dereferenced pointer")
	} else {
		fmt.Println("regular variable IS NOT equal to its dereferenced pointer")
	}

	if v2 == v3 {
		fmt.Println("two different pointers to the same variable are equal")
	} else {
		fmt.Println("two different pointers to the same variable are NOT equal")
	}
}
