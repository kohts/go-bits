package main

import (
	"flag"
	fmt "fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) GetName() string {
	return p.Name
}

func main() {
	var nFlag = flag.Int("n", 1234, "help message for flag n")
	configPath := flag.String("config", "config.yaml", "Filter")
	flag.Parse()

	fmt.Println("Hello, World!")
	fmt.Println("n has value: " + strconv.Itoa(*nFlag))
	fmt.Println("configPath has value: " + *configPath)

	aPerson := Person{
		Name: "Peter",
		Age:  43,
	}

	fmt.Println(aPerson.Name)
	fmt.Println(aPerson.GetName())
}
