package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func listen_and_dump(in_ch chan string, out_ch chan string) {
	for {
		out_ch <- "waiting"
		v := <-in_ch
		out_ch <- "running"

		fmt.Println(
			time.Now().Format("2006-01-02 15:04:05.000 MST: got a message"),
			v,
		)
	}
}

func dump(v interface{}, pfx interface{}) {
	switch pfx_typed := pfx.(type) {
	case string:
		fmt.Println(pfx_typed, ": ", v)
	default:
		fmt.Println(v)
	}
}

func amap() {
	var amap1 map[string]int
	var amap2 = map[string]string{
		"configuration": "main",
		"option1":       "value1",
	}

	if amap1 == nil {
		fmt.Println("nil map")
	}

	amap1 = make(map[string]int)
	dump(amap1, nil)

	amap1["test"] = 1
	dump(amap1, nil)

	amap1[strconv.Itoa(1)] = 1
	dump(amap1, nil)

	dump(amap1, nil)

	dump(amap2, "map literal")

	amap2["runtime_option"] = "test"
	dump(amap2, "map literal modified")

	c1, c2 := make(chan string), make(chan string)
	go listen_and_dump(c1, c2)

	c1 <- "test message"

	for <-c2 != "waiting" {
		fmt.Println("waiting for a child")
	}
}

func main2() {
	amap()
}

func main3() {
	naturals := make(chan int)
	squares := make(chan int)
	done := make(chan struct{})

	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": initialized channels")

	// counter
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": spawning counter")
	go func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": about to broadcast naturals")
		for x := 0; x <= 10; x++ {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": naturals <- ", x)
			naturals <- x
		}
		done <- struct{}{}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": done <- empty struct")
	}()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": spawned counter")

	// squarer
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": spawning squarer")
	go func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": about to read naturals")
		for x := range naturals {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": got from naturals", x)
			squares <- x * x
			fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": squares <- ", x*x)
		}
		close(squares)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": closed squares")
	}()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": spawned squarer")

	// printer
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": spawning printer")
	go func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": about to read squares")
		for x := range squares {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": got from squares: ", x)
			fmt.Println(x)
		}
	}()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": spawned printer")

	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": waiting <-done message")
	<-done
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": got <-done message, sending 90 to naturals")
	naturals <- 90
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": sent 90 to naturals")

	close(naturals)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000 MST"), ": closed naturals")

}

type Message struct {
	str     string
	wait    chan bool
	tooSlow bool
}

func boring(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			c <- Message{
				fmt.Sprintf("%s: %s %d\n", time.Now().Format("2006-01-02 15:04:05.000 MST"), msg, i),
				waitForIt,
				false,
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			//<-waitForIt
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func fanInSelect(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			case <-time.After(600 * time.Millisecond):
				c <- Message{
					fmt.Sprintf("%s: %s\n", time.Now().Format("2006-01-02 15:04:05.000 MST"), "no one said anything"),
					nil,
					true,
				}
			}
		}
	}()
	return c
}

func main() {
	//	c = fanIn(boring("Joe"), boring("Ann"))
	c := fanInSelect(boring("Joe"), boring("Ann"))

	for i := 0; i < 10; i++ {
		msg := <-c
		if !msg.tooSlow {
			fmt.Print(msg.str)
			// msg.wait <- true
		} else {
			fmt.Print(msg.str)
			break
		}
	}
	fmt.Println("You're boring; I'm leaving.")
}
