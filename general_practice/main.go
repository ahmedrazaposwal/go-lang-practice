package main

import (
	"fmt"
)

type Animal struct {
	name   string
	age    int
	legs   int
	weight float32
}

func myMessage() {
	fmt.Println("Print from function")
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := range 10 {
		if i == 5 {
			continue
		}
		if i == 7 {
			break
		}
		fmt.Println(i)
	}

	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "banana", "orange"}

	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	fruits_list := [3]string{"apple", "orange", "banana"}

	for idx, val := range fruits_list {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	for _, val := range fruits_list {
		fmt.Printf("%v\n", val)
	}

	myMessage()
	familyName("Ahmed")
	fmt.Printf("%d\n", add(5, 10, 15))
	fmt.Printf("%d\n", sum(5, 10))

	res := addition(10, 20)
	fmt.Println(res)

	fmt.Println(myFunction(5, "Hello"))
	a, b := myFunction(10, "Hello Go")

	fmt.Println(a, b)

	testCount(1)

	fmt.Println(factorial_resursion(4))

	var animal1 Animal

	animal1.name = "Lion"
	animal1.age = 10
	animal1.legs = 2
	animal1.weight = 70.0

	fmt.Println("Name: ", animal1.name)
	fmt.Println("Legs: ", animal1.legs)
	fmt.Println("Weight: ", animal1.weight)
	fmt.Println("Age: ", animal1.age)

	printAnimal(animal1)

}

func familyName(fname string) {
	fmt.Println("Hello", fname, "World!")
}

func add(x, y, z int) int {
	return x + y + z
}

func sum(x, y int) (result int) {
	result = x + y
	return result
}

func addition(x, y int) (result int) {
	result = x + y
	return
}

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func testCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testCount(x + 1)
}

func factorial_resursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_resursion(x-1)
	} else {
		y = 1
	}
	return y
}

func printAnimal(animal Animal) {
	fmt.Println("Name: ", animal.name)
	fmt.Println("Legs: ", animal.legs)
	fmt.Println("Weight: ", animal.weight)
	fmt.Println("Age: ", animal.age)
}
