package main

import "fmt"

// arg is 5? Factorials!
// 5 * (4 * (3 * (2 * (1 * 1)))) => 120 ?  lets see..
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n - 1)
}

func main() {
    fmt.Println("for 5:", fact(5)) // yup..
    fmt.Println("for 10:", fact(10))
}
