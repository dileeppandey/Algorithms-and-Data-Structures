// Find the largest number in an array

package main

import "errors"

func LargestNumber(numbers []int) (int, error) {
	if len(numbers) != 0 {
		largest := numbers[0]
		for _, val := range numbers {
			if largest < val {
				largest = val
			}
		}
		return largest, nil
	} else {
		return 0, errors.New("Array is empty!")
	}
}
