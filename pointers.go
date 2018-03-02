package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

// note the *
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1

    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // note the &, meaning its passing a pointer.
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // print the pointer => 0xc4200140c0
    fmt.Println("pointer:", &i)
}
