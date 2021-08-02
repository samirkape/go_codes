package main

import (
	"fmt"
	"strings"
)

func Reverse(str string) string {
	var s strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		s.WriteByte(str[i])
	}
	return s.String()
}

func Reverse1(s string) string {
	str := []byte(s)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}

func main() {

	input := "Samir"
	output := Reverse1(input)

	fmt.Println(output)
}
