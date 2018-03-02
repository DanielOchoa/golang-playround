package main

import "fmt"

func plus(a int, b int) int {
    return a + b
}

func plusPlus(a, b, c int) (int) {
    return a + b + c
}

// multiple return results

func vals() (int, int) {
    return 3, 7
}

func main() {
    // or res := plus(1, 2)
    var res int = plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)

    res1, res2 := vals()
    fmt.Println("multiple return vals:", res1, res2)

    _, c := vals()
    fmt.Println("just second val:", c)
}
