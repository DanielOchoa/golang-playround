package main

import "fmt"
import "time"

func main() {
    timer1 := time.NewTimer(time.Second * 2)


    <-timer1.C
    fmt.Println("timer 1 expired")

    // but if we want ability to cancel..

    timer2 := time.NewTimer(time.Second)

    go func() {
        <-timer2.C
        fmt.Println("timer 2 expired")
    }()
    stop2 := timer2.Stop()

    if stop2 {
        fmt.Println("timer 2 stoped")
    }
}