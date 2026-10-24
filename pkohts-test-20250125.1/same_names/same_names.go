package main

import "fmt"

type Command struct {
	Name      string
	Arguments string
}

func main() {
	Command := Command{
		Name:      "test",
		Arguments: "test",
	}
	fmt.Println(Command)

	fmt.Println("Hello, world!")
}
