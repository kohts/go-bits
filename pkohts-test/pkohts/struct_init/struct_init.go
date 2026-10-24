package main

import "fmt"

type Cat struct {
	name  string
	breed string
	age   int
	color string
}

type Speaker interface {
	Speak()
}

func main() {
	my_cats := [...]Cat{
		Cat{"Morpheus", "Tiffany", 10, "black"},
		Cat{"Patricia", "Hybrid", 5, "white"},
		Cat{"Shanti", "Hybrid", 5, "black and white"},
	}

	for _, c := range my_cats {
		fmt.Printf("%v\n", c)
	}

	var s Speaker
	fmt.Printf("%v\n", s)
}
