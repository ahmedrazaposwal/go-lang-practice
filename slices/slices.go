package main

import (
	"fmt"
)

func main() {
	var myslice = []int{}
	fmt.Println(myslice)

	my_slice := []int{1, 2, 3}
	fmt.Printf("slice: %v\n", my_slice)

	slice1 := []int{}
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice1)

	slice2 := []string{"Hello", "World!", "from", "Go"}
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println(slice2)

	arr1 := [...]int{1, 2, 3, 4, 5}
	slicefromArray := arr1[2:4]
	fmt.Printf("Slice: %v\n", slicefromArray)
	fmt.Printf("length: %d\n", len(slicefromArray))
	fmt.Printf("Capacity: %d\n", cap(slicefromArray))

	sliceByMake := make([]int, 5, 10)
	fmt.Printf("Slice: %v\n", sliceByMake)
	fmt.Printf("Length: %d\n", len(sliceByMake))
	fmt.Printf("Capacity: %d\n", cap(sliceByMake))

	sliceByMake1 := make([]int, 5)
	fmt.Printf("Slice: %v\n", sliceByMake1)
	fmt.Printf("Length: %d\n", len(sliceByMake1))
	fmt.Printf("Capacity: %d\n", cap(sliceByMake1))

	prices := []int{10, 20, 30}
	fmt.Println(prices[0])
	fmt.Println(prices)

	prices[2] = 50
	fmt.Println(prices)

	myslice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice:= %v\n", myslice1)
	fmt.Printf("Length:= %d\n", len(myslice1))
	fmt.Printf("Capacity:= %d\n", cap(myslice1))

	myslice1 = append(myslice1, 6, 7)
	fmt.Printf("Slice:= %v\n", myslice1)
	fmt.Printf("Length:= %d\n", len(myslice1))
	fmt.Printf("Capacity:= %d\n", cap(myslice1))

	myslice2 := []int{1, 2, 3}
	myslice3 := []int{4, 5, 6}
	myslice4 := append(myslice2, myslice3...)
	fmt.Println(myslice4)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Slice:= %v\n", numbers)
	fmt.Printf("Length:= %d\n", len(numbers))
	fmt.Printf("Capacity:= %d\n", cap(numbers))

	neededNumbers := numbers[:len(numbers)-3]
	fmt.Printf("Slice:= %v\n", neededNumbers)
	fmt.Printf("Length:= %d\n", len(neededNumbers))
	fmt.Printf("Capacity:= %d\n", cap(neededNumbers))

	newSlice := make([]int, len(neededNumbers))
	copy(newSlice, neededNumbers)

	fmt.Printf("Slice:= %v\n", newSlice)
	fmt.Printf("Length:= %d\n", len(newSlice))
	fmt.Printf("Capacity:= %d\n", cap(newSlice))

}
