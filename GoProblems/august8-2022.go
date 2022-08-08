// https://www.codewars.com/kata/

// Write a method, that will get an integer array as parameter and will process every number from this array.

// Return a new array with processing every number of the input-array like this:

// If the number has an integer square root, take this, otherwise square the number.

// Example
// [4,3,9,7,2,1] -> [2,9,3,49,4,1]
// Notes
// The input array will always contain only positive numbers, and will never be empty or null.



import "math"

func SquareOrSquareRoot(arr []int) []int {
  results := make([]int, len(arr))
  
  for i, x := range arr {
    sqrt := math.Sqrt(float64(x))
    
    if sqrt == math.Floor(sqrt) {
      results[i] = int(sqrt)
    } else {
      results[i] = x * x
    }
  }
  
  return results
}

//another solution

import (
	"math"
  )
  
  func SquareOrSquareRoot(arr []int) []int {
	var r []int
	for _, num := range arr {
	  s := math.Sqrt(float64(num))
	  if math.Mod(s, 1.0) == 0 {
		r = append(r, int(s))
	  } else {
		r = append(r, num*num)
	  }
	}
  
	return r
  }