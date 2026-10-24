package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := make([]int, 3)
	i := 0
	u := ""

	for u != "X" {
		fmt.Printf("Input [number or X to quit]: ")

		fmt.Scan(&u)

		if u == "X" {
			fmt.Printf("Exiting\n")
			break
		}

		v, err := strconv.Atoi(u)
		if err != nil {
			fmt.Printf("Not an integer: %v, retry\n", u)
			continue
		}

		if i < 3 {
			s[i] = v
		} else {
			s = append(s, v)
		}

		sort.Ints(s[0 : i+1])
		fmt.Printf("Array (sorted): %v\n", s[0:i+1])

		i = i + 1
	}
}
