package main

import "fmt"

func main() {

	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true // it also does inference if no type is specified
	var x = "dan"
	fmt.Println(d, x)

	var e int      // variables without value assigned are zero-valued.
	var e2 string  // it seems this may be just "".
	var e3 float32 // defaults to zero as well.
	fmt.Println(e, e2, e3)

	f := "short" // := shorthand for initializing and declaring var, e.g. var f string = "short"
	fmt.Println(f)
}
