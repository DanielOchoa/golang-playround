package main

import "fmt"

type rect struct {
    width, height int
}

// area method has a receiver of type *rect.
// references rect by pointer
func (r *rect) area() int {
    return r.width * r.height
}

// references rect by value (no pointer)
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area:", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area:", rp.area())
    fmt.Println("perim:", rp.perim())
}
