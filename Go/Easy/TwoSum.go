// Given an array and a number N, return true if two numbers in array sum to N,
// else return false

// Part 2: Return indices of two numbers

package main

import "fmt"

func TwoSum(numbers []int, target int) (bool, []int) {
	fmt.Printf("Find indices of numbers in %v that sum to %d\n", numbers, target)
	numbersMap := map[int]int{}
	for i, val := range numbers {
		j, ok := numbersMap[val]
		if ok {
			return true, []int{j, i}
		}
		numbersMap[target-val] = i
	}
	return false, []int{-1, -1}
}
