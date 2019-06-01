// string in go are immutable
// slicing str[start: end] -> the char at end index not included
// multi-line string can be constructed using ``
package main

func Reverse(str string) string {
	length := len(str)
	dest := []rune(str)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		dest[i], dest[j] = dest[j], dest[i]
	}
	return string(dest)
}
