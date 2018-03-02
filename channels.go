package main

import "fmt"

func main() {

    messages := make(chan string)

    go func() { messages<- "ping" }()

    msg := <-messages
    fmt.Println(msg)

    // buffered channels

    fmt.Println("buffering channels..")

    messages = make(chan string, 2)

    messages<- "buffered"
    messages<- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}

