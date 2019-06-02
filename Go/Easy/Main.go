package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's solve some easy level problems.")

	// Find the largest number in an array
	value, error := LargestNumber([]int{1, 4, 3, 900, 23})
	if error == nil {
		fmt.Println("Largest number in array [1, 4, 3, 900, 23] is:", value)
	}
	value, error = LargestNumber([]int{}) // passing an empty array
	if error == nil {
		fmt.Println("Largest number in array [1, 4, 3, 900, 23] is:", value)
	} else {
		fmt.Println(error)
	}

	// Two Sum problem
	fmt.Println(TwoSum([]int{1, 2, 3, 4, 5, 6, 7}, 7))
	fmt.Println(TwoSum([]int{1, 2, 3, 4, 5, 6, 7}, 54))
}
