package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read_stdio() string {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	return s
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }
type Bird struct{ name string }
type Snake struct{ name string }

func (c Cow) Eat()   { fmt.Printf("Cow %v ate some grass\n", c.name) }
func (c Cow) Move()  { fmt.Printf("Cow %v walked\n", c.name) }
func (c Cow) Speak() { fmt.Printf("Cow mooed: I'm %v\n", c.name) }

func (b Bird) Eat()   { fmt.Printf("Bird %v ate some worms\n", b.name) }
func (b Bird) Move()  { fmt.Printf("Bird %v flew\n", b.name) }
func (b Bird) Speak() { fmt.Printf("Bird peeped: I'm %v\n", b.name) }

func (s Snake) Eat()   { fmt.Printf("Snake %v ate some mice\n", s.name) }
func (s Snake) Move()  { fmt.Printf("Snake %v slithered\n", s.name) }
func (s Snake) Speak() { fmt.Printf("Snake hissed: I'm %v\n", s.name) }

func main() {
	my_animals := make(map[string]Animal)

	for {
		fmt.Printf("> ")
		input_str := strings.TrimSpace(read_stdio())

		input_arr := strings.Split(input_str, " ")
		if len(input_arr) != 3 {
			fmt.Printf("wrong input format, should be 3 space separated values\n")
			continue
		}

		if input_arr[0] == "newanimal" {
			_, exists := my_animals[input_arr[1]]
			if exists {
				fmt.Printf("There's already an animal with that name!\n")
				continue
			}

			var a Animal
			switch input_arr[2] {
			case "cow":
				a = Cow{input_arr[1]}
			case "bird":
				a = Bird{input_arr[1]}
			case "snake":
				a = Snake{input_arr[1]}
			default:
				fmt.Printf("Unknown animal type %v\n", input_arr[2])
				continue
			}
			my_animals[input_arr[1]] = a
			fmt.Printf("Created it!\n")
		} else if input_arr[0] == "query" {
			a, exists := my_animals[input_arr[1]]
			if !exists {
				fmt.Printf("No animal by such name: %v\n", input_arr[1])
				continue
			}

			switch input_arr[2] {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			default:
				fmt.Printf("Unknown action: %v\n", input_arr[2])
				continue
			}
		} else {
			fmt.Printf("Unknown command: %v\n", input_arr[0])
			continue
		}
	}
}
