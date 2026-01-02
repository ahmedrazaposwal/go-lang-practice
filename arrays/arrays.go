package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2}
	fmt.Println(arr1)

	var arr2 = [...]string{"Go"}
	fmt.Println(arr2)

	var arr3 = [2]string{"Go", "Hello"}
	fmt.Println(arr3)

	arr4 := [2]int{1, 2}
	arr5 := [...]int{1, 2, 3}

	fmt.Println(arr4, arr5)
	fmt.Println(arr3[1])

	arr3[0] = "Hello"
	arr3[1] = "World!"
	fmt.Println(arr3)

	arr6 := [5]int{}
	fmt.Println(arr6)

	arr7 := [10]int{1: 5, 3: 50, 9: 100}
	fmt.Println(arr7)
	fmt.Println(len(arr7))
	arr8 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(arr8))

}
