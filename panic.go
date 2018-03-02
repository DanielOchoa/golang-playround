package main

import "os"

func main() {
    panic("a problem")

    // it works like a return value, this will never be reached.
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
