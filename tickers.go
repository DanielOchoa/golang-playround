package main
import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(time.Millisecond * 500)

    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

    // why use sleep lol?
    // time.Sleep(time.Millisecond * 1600)

    timer := time.NewTimer(time.Millisecond * 1600)

    <-timer.C
    ticker.Stop()

    fmt.Println("Ticker stopped")
}

