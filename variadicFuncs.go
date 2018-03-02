
package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    var total int = 0
    // with for loop (good exercise, but range is more succint)
    for idx, length := 0, len(nums); idx < length; idx++ {
        total += nums[idx]
    }
    /*for _, num := range nums {
        total += num
    }*/

    fmt.Println(total)
}

func main() {
    sum(1, 2, 3, 4, 5, 6)
    sum(1, 3, 3)

    // SPREAD Op!
    // nums := make([]int, 4) // var nums[5] int <- doesn't work since nums[5] is diff type than nums[] ... go figure.
    // for i := 0; i < 5; i++ {
        //nums = append(nums, i + 1) // <- only works on slices.
    //    nums = append(nums, i + 1)
    // }

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
