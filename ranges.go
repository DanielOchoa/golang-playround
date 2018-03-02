package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}

    sum := 0

    // The `_` is the index.
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // index print
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // for maps, its key - val
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for key, value := range kvs {
        fmt.Printf("%s -> %s\n", key, value)
    }

    // print just key
    for key := range kvs {
        fmt.Printf("%s\n", key)
    }

    // range on strings iterates over Unicode code points. The first value is the starting byte
    // index of the `rune` and the second the `rune` itself.
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
