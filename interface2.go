package main

import "fmt"

// weirdness in interfaces.. every type implements at LEAST an empty interface...
// that is, `interface{}`. So if a func takes an empty interface as param, that means
// it would accept any param type.

func doSomething(v interface{}) {
    fmt.Printf("val: %v, type: %T\n", v, v)
}

func main() {
    doSomething("hi")
    doSomething(2)
    doSomething(2.3)
    doSomething([]int{1, 2, 3})
}
