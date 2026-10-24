package main

import (
    "fmt"
    "example/greetings"
)

func main() {
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
