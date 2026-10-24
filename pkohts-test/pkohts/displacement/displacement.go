package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_stdio() string {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	return s
}

func read_stdio_float() float64 {
	s := strings.TrimSpace(read_stdio())

	s1, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("not a float in the string: %v\n", s)
	}

	return s1
}

func GenDisplaceFn(a, v0, s0 float64) func(int) float64 {
	return func(elapsed_time int) float64 {
		t := float64(elapsed_time)
		s := 0.5*a*t*t + v0*t + s0
		return s
	}
}

func main() {
	fmt.Printf("Please enter acceleration (float): ")
	acceleration := read_stdio_float()

	fmt.Printf("Please enter initial velocity (float): ")
	initial_velocity := read_stdio_float()

	fmt.Printf("Please enter initial displacement (float): ")
	initial_displacement := read_stdio_float()

	fmt.Printf("Please enter time (integer): ")
	time_str := strings.TrimSpace(read_stdio())
	time, err := strconv.Atoi(time_str)
	if err != nil {
		log.Fatalf("not an integer in the string: %v\n", time_str)
	}

	fmt.Printf("Input parameters: %v, %v, %v, %v\n", acceleration, initial_velocity, initial_displacement, time)

	fn := GenDisplaceFn(acceleration, initial_velocity, initial_displacement)
	fmt.Printf("Displacement after %v second(s): %v\n", time, fn(time))
}
