// problem URL https://www.codewars.com/kata/59cfc000aeb2844d16000075

// DESCRIPTION:
// Given a string, capitalize the letters that occupy even indexes and odd indexes separately, and return as shown below. Index 0 will be considered even.

// For example, capitalize("abcdef") = ['AbCdEf', 'aBcDeF']. See test cases for more examples.

// The input will be a lowercase string with no spaces.

// Good luck!

// My solution

package kata

import "strings"

func Capitalize(st string) []string {
	var answer = make([]string, 2)
	for i, _ := range st {
		if i%2 == 0 {
			answer[0] += strings.ToUpper(string(st[i]))
		} else {
			answer[0] += strings.ToLower(string(st[i]))
		}
	}

	for i, _ := range st {
		if i%2 == 1 {
			answer[1] += strings.ToUpper(string(st[i]))
		} else {
			answer[1] += strings.ToLower(string(st[i]))
		}
	}
	return answer
}

// I liked that I demonstrated the use of `for` loops in a `range`.  I used `make` to make a `slice.`

// I learned about how to use the `strings.ToUpper` and `strings.ToLower` methods.  It converts strings to their upper or lowercase versions.  `st[i]` in these contexts are runes. It would have been better just to use their `values` instead of their indexes because of the complexity of working with runes.

//An answer I liked and learned from:

package kata
import "unicode"
func Capitalize(s string) []string {
  a, b := []rune(s),[]rune(s)
  for i := range a {
    if i%2 == 0 {
      a[i] = unicode.ToUpper(a[i])
    }else{
      b[i] = unicode.ToUpper(b[i])
    }
  }
  return []string{string(a), string(b)}
}

//This person used the `unicode` package.  They initialized `a` and `b` to slices of runes with string `s` as the argument.

// `unicode.ToUpper` converts the rune directly to a letter.

//Another answer I liked

package kata

import "unicode"

func Capitalize(st string) []string {
  var st1, st2 string
  p1 := &st2
  p2 := &st1
  for _, c := range st {
    p1, p2 = p2, p1
    *p1 += string(unicode.ToUpper(c))
    *p2 += string(unicode.ToLower(c))
  }
  return []string{st1, st2}
}

//This person did multiple assignment and is always swapping the pointers.  