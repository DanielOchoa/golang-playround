package main

import "fmt"
import "sort"

func main() {
    strs := []string{"c", "a", "b"}

    sort.Strings(strs)

    fmt.Println("Strings:", strs)

    // now ints
    ints := []int{7, 2, 4}
    sort.Ints(ints)

    fmt.Println("Ints:", ints)

    // check if already sorted
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted:", s)
}
