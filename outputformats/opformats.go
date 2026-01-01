package main

import (
	"fmt"
)

func main() {
	var i, j = "Hello Go", 4
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v %T\n", j, j)
	fmt.Printf("%#v\n", j)

	fmt.Printf("%v%%\n", 15)

	var n = 15

	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%+d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%O\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%#x\n", n)
	fmt.Printf("%#X\n", n)
	fmt.Printf("%5d\n", n)
	fmt.Printf("%-5d\n", n)
	fmt.Printf("%+5d\n", n)
	fmt.Printf("%05d\n", n)
	fmt.Printf("%-+5d\n", n)

	var str = "HelloGo"

	fmt.Printf("%s\n", str)
	fmt.Printf("%q\n", str)
	fmt.Printf("%8s\n", str)
	fmt.Printf("%-7s\n", str)
	fmt.Printf("%x\n", str)
	fmt.Printf("% x\n", str)

	var k = true
	fmt.Printf("%t\n", k)

	var l = 3.141

	fmt.Printf("%e\n", l)
	fmt.Printf("%f\n", l)
	fmt.Printf("%.3f\n", l)
	fmt.Printf("%5.2f\n", l)
	fmt.Printf("%g\n", l)

}
