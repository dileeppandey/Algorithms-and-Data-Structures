// A number is an even ended number it its start and end digits are same.

package main

import "fmt"

func EvenEndedNumber(number int) bool {
	s := fmt.Sprintf("%d", number)
	return s[0] == s[len(s)-1]
}
