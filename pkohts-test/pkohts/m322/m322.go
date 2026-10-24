package main

import "fmt"

func main() {
	type Config struct {
		config  map[string]string
		version int
	}

	var aCfg = new(Config)
	aCfg.config["best"] = "test"
	fmt.Printf("%v\n", aCfg)

	s := make([]int, 5)
	s[0] = 1
	fmt.Printf("%v\n", s)

	//c := make(map[string]string)
	c := map[string]string{
		"test": "best",
		"zest": "lest",
	}

	c["test"] = "best"
	c["zest"] = "lest"
	fmt.Printf("%v\n", c)
}
