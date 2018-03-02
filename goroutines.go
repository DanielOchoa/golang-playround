package main

import "fmt"

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {
    // sync function call
    f("direct")

    // start a thread? goroutines!!! will execute concurrently.
    go f("goroutine")

    // anonym func goroutine too
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    fmt.Println("All functions have been called...")
    // our goroutines run separately, so execution falls through.
    // The `Scanln` code requires we press a key before program exits.
    var input string
    fmt.Scanln(&input)
    fmt.Println("done!")
}
