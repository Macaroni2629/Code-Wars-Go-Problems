// https://www.codewars.com/kata/57a0e5c372292dd76d000d7e

// Write a function that accepts an integer n and a string s as parameters, and returns a string of s repeated exactly n times.

// Examples (input -> output)
// 6, "I"     -> "IIIIII"
// 5, "Hello" -> "HelloHelloHelloHelloHello"

//Best answer

import "strings"

func RepeatStr(repititions int, value string) string {
	return strings.Repeat(value, repititions)
}

//Another good answer

func RepeatStr(repititions int, value string) (result string) {
	for i := 0; i < repititions; i++ {
		result += value
	}
	return
}