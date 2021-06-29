package main

import (
	"fmt"
)

func abc(c chan string) {
	c <- "Hello - abc"
	c <- "Goroutine - abc"
	c <- "Chanel -abc"
}
func def(c chan string) {
	c <- "Hello - def"
	c <- "GoLang - def"
	c <- "Perfect - def"
}
func main() {
	var channel1 = make(chan string)
	var channel2 = make(chan string)
	go abc(channel1)
	go def(channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)

}
