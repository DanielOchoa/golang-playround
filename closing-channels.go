package main

import "fmt"

func main() {
    // This would say "all goroutines are asleep - deadlock!"
    // probably from waiting that last Println when we only sent 3 messages
    // and not four.
    /*
    mychan := make(chan string)

    go func() {
        mychan<- "hi"
        mychan<- "ho"
        mychan<- "hi"
    }()

    fmt.Println(<-mychan)
    fmt.Println(<-mychan)
    fmt.Println(<-mychan)
    fmt.Println(<-mychan)
    */

    //// closing channels seems to help..
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            // more will be false when jobs chan is closed and all values have been received
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done<- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs<- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // with synchronization approach, program runs until done receives a value, so we can send it.
    <-done
}
