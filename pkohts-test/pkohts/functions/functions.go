package main

import "fmt"

/*func fA(fB (int) string) func (string) int {
	fmt.Printf("\n")
}*/

/*func fA(fB func(int) string) {
	fmt.Printf("\n")
}*/

/*func fA(int) string {
	return ""
}*/

/*
func fA() fB func(string) int {
	fmt.Printf("\n")
}
*/

func fA(fB func(string) int) fB func(string) int {
	return func(v string) int { return fB("test") + len(v) }
}

func main() {
	fmt.Printf("\n")

	f1 := fA(func(v string) int { return len(v) })

	fmt.Printf("%v\n", f1("test"))
	fmt.Printf("%v\n", f1("bright"))
}
