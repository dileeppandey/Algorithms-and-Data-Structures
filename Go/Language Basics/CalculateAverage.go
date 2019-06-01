// Numbers and Assignments

// := automatically infers the type of a variable and assigns the value, used
// ony for at least one new variable declaration

// Go has strict type checking, cannot do operations on int and float

// It is convention to capitalize first letter of public functions

package main

func CalculateAverage(nums ...int) float32 {
	var sum int
	sum = 0
	for _, value := range nums {
		sum += value
	}
	average := float32(sum) / float32(len(nums))
	return average
}
