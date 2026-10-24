package channels

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%s: before creating channel\n", time.Now().Format("2006-01-02 15:04:05.000 MST"))
	ch1 := make(chan string, 1)
	fmt.Printf("%s: after creating channel\n", time.Now().Format("2006-01-02 15:04:05.000 MST"))
	ch1 <- "sending a message to a channel"
	fmt.Printf("%s: after sending a message to the channel\n", time.Now().Format("2006-01-02 15:04:05.000 MST"))
	val := <-ch1
	fmt.Printf("%s: after reading a message from the channel: [%s]\n", time.Now().Format("2006-01-02 15:04:05.000 MST"), val)
}
