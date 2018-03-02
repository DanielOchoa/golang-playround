package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan string)

    // non-blocking receive
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
        break
    default:
        fmt.Println("no message received...")
    }

    // non-blocking send
    var msg string = "hi"

    select {
    case messages<- msg:
        fmt.Println("sent msg", msg)
    default:
        fmt.Println("no message sent...")
    }

    // on multiple channels, receives
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
