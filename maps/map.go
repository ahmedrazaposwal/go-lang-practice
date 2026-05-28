package main

import (
	"fmt"
)

func main() {
	var map1 = map[string]string{"brand": "Toyota", "Model": "Auto", "year": "1996"}
	map2 := map[string]int{"Oslo": 1, "Bergen": 2, "Stavanger": 3}

	fmt.Println(map1)
	fmt.Println(map2)

	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["Model"] = "Mustang"
	a["year"] = "2020"

	b := make(map[string]int)
	b["a"] = 1
	b["b"] = 2

	fmt.Println(a)
	fmt.Println(b)

	var c map[string]int
	var d = make(map[string]string)

	fmt.Println(c == nil)
	fmt.Println(d == nil)

	fmt.Println(a["year"])
	a["year"] = "2025"
	a["color"] = "red"
	fmt.Println(a)

	delete(a, "color")
	fmt.Println(a)

	_, ok := a["year"]
	fmt.Println(ok)

	_, ok1 := a["day"]
	fmt.Println(ok1)

	e := map[string]string{"brand": "Ford", "color": "red", "year": "2025"}
	f := e

	fmt.Println(e)
	fmt.Println(f)

	e["color"] = "green"
	fmt.Println("After change to e:")

	fmt.Println(e)
	fmt.Println(f)

	for k, v := range e {
		fmt.Printf("%v : %v, ", k, v)
	}

	g := map[string]int{"one": 1, "two": 2, "three": 3}
	var h []string
	h = append(h, "one", "two", "three")

	fmt.Println()

	for k, v := range g {
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range h {
		fmt.Printf("%v : %v, ", element, g[element])
	}

	fmt.Println()

	i := 1
	for i < 3 {
		fmt.Println("hello")
		i++
	}

	for {
		fmt.Println("loop")
		break
	}

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())

	fmt.Println(fact(5))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

	for u, w := range "abc" {
		fmt.Println(u, w)
	}

	z := 5
	fmt.Println(z)

	funByValue(z)
	fmt.Println(z)

	funByReference(&z)
	fmt.Println(z)
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, val := range nums {
		total += val
	}
	fmt.Println(total)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func funByValue(n int) {
	n = 0
}

func funByReference(iptr *int) {
	*iptr = 0
}
