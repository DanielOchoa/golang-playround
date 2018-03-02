package main

import "fmt"

// only accepts a channel for SENDING values.
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// accepts a channel for RECEIVING values, and another for SENDING.
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)

    go ping(pings, "passed message")
    go pong(pings, pongs)
    fmt.Println(<-pongs)
}
