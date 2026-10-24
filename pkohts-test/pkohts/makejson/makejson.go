package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var t string

	p := map[string]string{}

	fmt.Printf("Input person name: ")
	fmt.Scanln(&t)
	p["name"] = t
	fmt.Printf("Input person address: ")
	fmt.Scan(&t)
	p["address"] = t

	j, e := json.Marshal(p)
	if e != nil {
		fmt.Printf("Error marshaling p: %v", e)
		os.Exit(1)
	}
	fmt.Printf("%v\n", string(j))
}
