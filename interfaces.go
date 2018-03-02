package main

import "fmt"
import "math"

// define interface
type geometry interface {
    area() float64
    perim() float64
}

// structs rect and circle...
type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

// implement interface on rect and circle...
// rect is the receiver..
func (r rect) area() float64 {
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * math.Pow(2, c.radius)
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// generic `measure` func, var that uses can call interface
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

/**
 * Anima example now
 * Interfaces are designs of what something can do, not what it holds.
 */
type Animal interface {
    speak() string
}

type Dog struct {
}

func (d Dog) speak() string {
    return "Woof!"
}

type Cat struct {}
func (c *Cat) speak() string {
    return "Meooww!"
}

type Llama struct {}
func (l Llama) speak() string {
    return "????"
}

type JavaProgrammer struct {}
func (p JavaProgrammer) speak() string {
    return "DesignPatterns!"
}

/* end */

func main() {
    f := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(f)
    measure(c)

    // our contrived example
    // simple..
    d := Dog{}
    fmt.Println(d.speak())
    // but.. notice how the slice type is interface, but takes our types.
    // this fails if you don't pass pointer receiver to Cat
    // using new() is same as & prefix.

    // We can use &Dog, even thought it is *Dog and not Dog who implements animal interface - because &Dog
    // can get the value, through the pointer, but not the other way around - that's why we can't use Cat{}
    // because multiple pointers (thus possible multiple values) could be references Cat{}.
    animals := []Animal{new(Dog), &Cat{}, Llama{}, JavaProgrammer{}}

    for _, animal := range animals {
        fmt.Println(animal.speak())
    }
}
