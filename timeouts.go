package main

import (
    "time"
    "fmt"
)

func main() {
    chan1 := make(chan string, 1)

    go func() {
        time.Sleep(time.Second * 2)
        chan1<- "call 1!"
    }()

    select {
    case res := <-chan1:
        // this dones't run cuz func sleep is 2
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout of call 1")
    }

    chan2 := make(chan string, 1)

    go func() {
        time.Sleep(time.Second * 3)
        chan2<- "result 2"
    }()

    select {
    case res := <-chan2:
        fmt.Println(res)
    case <-time.After(time.Second * 2):
        fmt.Println("timeout 2")
    }
}
