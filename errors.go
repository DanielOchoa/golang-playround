package main

import (
    "fmt"
    "errors"
)

func f1(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("Can't work with 42")
    }
    return arg + 3, nil
}

type argError struct {
    arg int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "Can't work with it"}
    }
    return arg + 3, nil
}

func main() {
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }

    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    // if you want to programatically use the error data in a custom way, you'll need to get the error
    // as an instance of the custom error type via assertion..
    _, e := f1(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
