package main

import "fmt"

// struct
type person struct {
    name string
    age  int
}

func main() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"})

    // pointer
    fmt.Println(&person{name: "Ann", age: 40})

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(s.name)

    sp.age = 51
    fmt.Println(sp.age)
    fmt.Println(s.age)

    s.age = 20
    fmt.Println(s.age)
    fmt.Println(sp.age)
}
