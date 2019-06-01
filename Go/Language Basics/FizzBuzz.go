// Go only has for loops
// Here for is implemented as a while loop construct
// Go has two conditional construct - if and switch
// switch can be used as naked switch like below or regular switch construct

package main

import "fmt"

func FizzBuzz(number int) {
	val := 1
	for val <= number {
		switch {
		case val%3 == 0 && val%5 == 0:
			fmt.Println(val, "FizzBuzz")
		case val%3 == 0:
			fmt.Println(val, "Fizz")
		case val%5 == 0:
			fmt.Println(val, "Buzz")
		}
		val += 1
	}

}
