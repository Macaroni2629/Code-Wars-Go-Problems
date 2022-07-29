// https://www.codewars.com/kata/56cd44e1aa4ac7879200010b/train/go

// Is the string uppercase?
// Task
// Create a method to see whether the string is ALL CAPS.

// Examples (input -> output)
// "c" -> False
// "C" -> True
// "hello I AM DONALD" -> False
// "HELLO I AM DONALD" -> True
// "ACSKLDFJSgSKLDFJSKLDFJ" -> False
// "ACSKLDFJSGSKLDFJSKLDFJ" -> True
// In this Kata, a string is said to be in ALL CAPS whenever it does not contain any lowercase letter so any string containing no letters at all is trivially considered to be in ALL CAPS.

// I could not come up with an answer for this one.

//Best solution from CodeWars

package kata

import (
  "strings"
)

type MyString string

func (s MyString) IsUpperCase() bool {
  return s == MyString(strings.ToUpper(string(s)))
}

//This solution is good because it is so short but concise.  It uses the `strings.ToUpper` method to uppercase `s` myString type.  Note that the developer had to convet the `MyString` type to `myString`.

package kata
import (
  "unicode"
)
type MyString string

func (s MyString) IsUpperCase() bool {
  for _, char := range s {
    if unicode.IsLower(char) {
      return false
    }
  }
  
  return true
}

//This solution is good for practicing iteration.  

//It seems like with white spaces, white spaces are read as uppercase in Golang.
