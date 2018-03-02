package main

import "fmt"

func main() {
    // map([key-type]val-type)
    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1:", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("removed 'k2':", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // "" can be a key, or zero, or false.
    m[""] = 1111
    _, prs2 := m[""]
    fmt.Println("prs2:", prs2)

    // map with only two possible keys!
    falsy := make(map[bool]string)
    falsy[false] = "FALSE"
    falsy[true] = "TRUE"

    falsyVal, present := falsy[true]
    fmt.Println("falsy:", falsyVal, present)


    // declare AND init in one line
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("inited map:", n)

    test = map[bool]string{false: "1"}
    fmt.Println(test)

}
