package main

import "fmt"

func main2() {
	x := 7
	switch {
	case x > 3:
		fmt.Printf("1")
	case x > 5:
		fmt.Printf("2")
	case x == 7:
		fmt.Printf("3")
	default:
		fmt.Printf("4")
	}
}

func main3() {
	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)
}

func main4() {
	var x int
	var y *int
	z := 3
	y = &z
	x = *y + 2
	fmt.Printf("%v\n", x)
}

func main5() {
	x := []int{4, 8, 5}
	y := -1
	for _, elt := range x {
		if elt > y {
			y = elt
		}
	}
	fmt.Print(y)
}

func main6() {
	x := [...]int{4, 8, 5}
	y := x[0:2]
	z := x[1:3]
	y[0] = 1
	z[1] = 3
	fmt.Print(x)
}

func main7() {
	x := [...]int{1, 2, 3, 4, 5}
	y := x[0:2]
	z := x[1:4]
	fmt.Print(len(y), cap(y), len(z), cap(z))
}

func main8() {
	x := map[string]int{
		"ian": 1, "harris": 2}
	for i, j := range x {
		if i == "harris" {
			fmt.Print(i, j)
		}
	}
}

type P struct {
	x string
	y int
}

func main9() {
	b := P{"x", -1}
	a := [...]P{P{"a", 10}, P{"b", 2}, P{"c", 3}}
	for _, z := range a {
		if z.y > b.y {
			b = z
		}
	}
	fmt.Println(b.x)
}

func main() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
}
