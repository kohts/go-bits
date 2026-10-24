package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Swap(aNumbers []int, swapTo int) {
	if swapTo+1 >= len(aNumbers) {
		log.Fatalf("Swap got invalid slice position: %v (max allowed is %v)\n", swapTo, len(aNumbers)-1)
	}

	tmp := aNumbers[swapTo]
	aNumbers[swapTo] = aNumbers[swapTo+1]
	aNumbers[swapTo+1] = tmp
}

func BubbleSort(aNumbers []int) []int {
	if len(aNumbers) < 2 {
		return aNumbers
	}

	swappedSomething := true

	for swappedSomething {
		swappedSomething = false

		for i := 0; i+1 < len(aNumbers); i++ {
			if aNumbers[i] > aNumbers[i+1] {
				Swap(aNumbers, i)
				swappedSomething = true
			}
		}
	}

	return aNumbers
}

func main() {
	var s string
	fmt.Printf("Please enter a string of up to 10 space separated integers: ")

	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	s = strings.TrimSpace(s)

	intSlice := []int{}
	for _, s = range strings.Split(s, " ") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Not an integer in the string: %v, \n", s)
		}

		intSlice = append(intSlice, i)
	}

	if len(intSlice) > 10 {
		log.Fatalf("not more than 10 integers allowed, got %v\n", len(intSlice))
	}

	fmt.Printf("Converted slice with integers: %v\n", intSlice)
	sortedNumbers := BubbleSort(intSlice)
	fmt.Printf("Sorted slice with integers: %v\n", sortedNumbers)
}
