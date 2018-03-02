package main

import "fmt"

func main() {
    queue := make(chan string, 2)

    queue<- "one"
    queue<- "two"

    close(queue)

    // channel is closed but we still get to receive elements.
    for elem := range queue {
        fmt.Println(elem)
    }
}
