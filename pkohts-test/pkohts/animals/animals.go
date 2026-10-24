package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func read_stdio() string {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	return s
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {
		fmt.Printf("> ")
		input_str := strings.TrimSpace(read_stdio())
		input_arr := strings.Split(input_str, " ")
		if len(input_arr) != 2 {
			fmt.Printf("wrong input format, should be two space separated values\n")
			continue
		}

		var obj Animal
		if input_arr[0] == "cow" {
			obj = cow
		} else if input_arr[0] == "bird" {
			obj = bird
		} else if input_arr[0] == "snake" {
			obj = snake
		} else {
			fmt.Printf("unknown creature: %v\n", input_arr[0])
			continue
		}

		if input_arr[1] == "eat" {
			obj.Eat()
		} else if input_arr[1] == "move" {
			obj.Move()
		} else if input_arr[1] == "speak" {
			obj.Speak()
		} else {
			fmt.Printf("unknown action: %v\n", input_arr[1])
			continue
		}
	}
}
