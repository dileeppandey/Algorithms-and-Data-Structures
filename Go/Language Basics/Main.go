// TODO: Replace function call with unit tests.

package main

import (
	"fmt"
)

func calculateAverage() {
	fmt.Println("Let's calculate average!")
	fmt.Println("Average of 1,2,3,4 is", CalculateAverage(1, 2, 3, 4))
}

func calculateWage() {
	fmt.Println("Let's calculate wage of workers!")
	fmt.Println("A worked 40 hrs, $15/hr standard & $20 overtime", CalculateWage(40.0, 15.0, 20.0))
	fmt.Println("A worked 55 hrs, $15/hr standard & $20 overtime", CalculateWage(55.0, 15.0, 20.0))
}

func fizzBuzz() {
	fmt.Println("Let's solve FizzBuzz")
	FizzBuzz(20)
}

func reverseString() {
	fmt.Println("Let's reverse a 'string' ->", Reverse("string"))
	fmt.Println("Let's reverse a string 'ababab' ->", Reverse("ababab"))
}

func evenEndedNumber() {
	fmt.Println("Let's check to see if 1 is an even ended number", EvenEndedNumber(1))
	fmt.Println("Let's check to see if 121 is an even ended number", EvenEndedNumber(121))
	fmt.Println("Let's check to see if 12 is an even ended number", EvenEndedNumber(12))
}

func wordCount() {
	text := `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
	tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
	quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
	consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
	cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
	proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

	fmt.Println("Count of words in Lorem Ipsum: ", WordCount(text))
}

func square() {
	square, err := NewSquare(0, 0, 4)
	if err == nil {
		fmt.Println("Center of Square is:", square.Center)
		fmt.Println("Length of Square is:", square.Length)
		fmt.Println("Area of Square is:", square.Area())
		fmt.Println("Move Center of Square")
		square.Move(1,1)
		fmt.Println("New Center of Square is:", square.Center)
	}
}

func main() {
	fmt.Println("Welcome to Go!")
	// calculateAverage()
	// calculateWage()
	// fizzBuzz()
	// reverseString()
	// evenEndedNumber()
	// wordCount()
	square()
}
