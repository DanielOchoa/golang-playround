package main

import "fmt"
import "os"

type point struct {
    x, y int
}

func main() {

    comChan := make(chan point)


    go func() {
        // print struct
        p := point{1, 2}
        comChan<- p
    }()

    p := <-comChan

    // print values
    fmt.Printf("%v\n", p)

    // print with inside var names
    fmt.Printf("%+v\n", p)

    // print with Go syntax representation of the value, e.g. `main.point{x1, y:2}`
    fmt.Printf("%#v\n", p)

    // print type
    fmt.Printf("%T\n", p)

    // boolean formatting
    fmt.Printf("%t\n", true)

    // print number... ?
    fmt.Printf("%d\n", 123)

    // binary representation
    fmt.Printf("%b", 14)

    // corresponding to any given integer
    fmt.Printf("%c\n", 123)

    // hex // 1c8
    fmt.Printf("%x\n", 456)

    // float formatting `78.00000`
    fmt.Printf("%f\n", 78.6)

    // scientific notations ...e+08
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)

    // string print
    fmt.Printf("%s\n", "\"string\"")

    // double quote strings
    fmt.Printf("%q\n", "\"string\"")

    // more hex
    fmt.Printf("%x\n", "hex this")

    // pointer representation
    fmt.Printf("%p\n", &p)

    // filter numbers to places, e.g. float them right
    fmt.Printf("|%6d|%6d|\n", 12, 345)
    // 6 is the width of the box where number is, and after . it's precission.
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
    // left justified
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
    // aligning strings, use for tables.
    fmt.Printf("|%6s|%6s|\n", "foo", "b")
    // left justify
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    // save into a variable instead of stdout
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)

    // You can format+print to io.Writers other than os.Stdout using Fprintf.
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
