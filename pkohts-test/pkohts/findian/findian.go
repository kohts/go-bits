package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	fmt.Printf("Please enter a string: ")

	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	s = strings.TrimSpace(s)

	if len(s) > 2 &&
		strings.ToLower(string(s[0])) == "i" &&
		strings.ToLower(string(s[len(s)-1])) == "n" &&
		strings.Contains(strings.ToLower(s), "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
