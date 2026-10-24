package main

import "fmt"

func main() {
	m1 := make(map[string]string)

	m1["test"] = "test"

	switch {
	case m1["best"] == "":
		fmt.Println("non-existing map element is an empty string")
	default:
		fmt.Println("non-existing map element is none of nil or empty, what is it then?")
	}

	fmt.Printf("-%s-", m1["best"])
}
