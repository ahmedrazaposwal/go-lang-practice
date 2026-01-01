package main

import (
	"fmt"
)

var x, y int = 10, 20

func main() {
	var a, b int = 1, 4
	fmt.Println(a, b)

	var c, d = 5, 6
	fmt.Println("c+d=", c+d)

	var str = "Go"
	var str1 string = "lang"
	fmt.Println(str + str1)

	x := 5
	fmt.Println(x)

	fmt.Println(y, "/", x, "=", y/x)

	y = 30
	fmt.Println(y, "/", x, "=", y/x)

	var t = false
	fmt.Println(t)

	var u bool = true
	fmt.Println(u)

	s := "Go learning"
	fmt.Println(s)

	var m, n = 6, "Hello"
	o, p := 7, "Go"
	fmt.Println(m, n, o, p)

	var (
		q int
		r int    = 1
		w string = "Blue Sky"
	)
	fmt.Println(q, r, w)

}
