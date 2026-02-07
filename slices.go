package main

import "fmt"

func slices() {
	// Create slice
	numbers := []int{1, 2, 3, 4, 5}

	// Empty slice
	var empty []int
	empty = []int{}

	// Make slice with capacity
	sliceLen := make([]int, 5)     // length 5, cap 5
	sliceCap := make([]int, 5, 10) // length 5, cap 10

	// Append elements
	numbers = append(numbers, 6)
	numbers = append(numbers, 7, 8, 9)

	// Access elements
	first := numbers[0]
	last := numbers[len(numbers)-1]

	// Slicing
	subset1 := numbers[1:4] 
	subset2 := numbers[:3]
	subset3 := numbers[2:]  
	subset4 := numbers[:]  

	// Length and capacity
	length := len(numbers) // Number of elements
	capacity := cap(numbers)

	// Loop through slice
	for i, num := range numbers {
		fmt.Println(i, num)
	}

	// Use variables to avoid unused errors
	fmt.Println(empty, sliceLen, sliceCap, first, last, subset1, subset2, subset3, subset4, length, capacity)
}
