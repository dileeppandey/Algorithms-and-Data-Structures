package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go!")
	fmt.Println("Let's calculate average!")
	fmt.Println("Average of 1,2,3,4 is", CalculateAverage(1, 2, 3, 4))

	fmt.Println("Let's calculate wage of workers!")
	fmt.Println("A worked 40 hrs, $15/hr standard & $20 overtime", CalculateWage(40.0, 15.0, 20.0))
	fmt.Println("A worked 55 hrs, $15/hr standard & $20 overtime", CalculateWage(55.0, 15.0, 20.0))

	fmt.Println("Let's solve FizzBuzz")
	FizzBuzz(20)
}