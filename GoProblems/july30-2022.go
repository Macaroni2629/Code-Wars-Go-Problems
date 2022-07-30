// https://www.codewars.com/kata/5513795bd3fafb56c200049e/train/go

// Create a function with two arguments that will return an array of the first (n) multiples of (x).

// Assume both the given number and the number of times to count will be positive numbers greater than 0.

// Return the results as an array (or list in Python, Haskell or Elixir).

// Examples:

// countBy(1,10)  should return  {1,2,3,4,5,6,7,8,9,10}
// countBy(2,5)  should return {2,4,6,8,10}

//Good answer
package kata

func CountBy(x, n int) []int {
	arr := []int{}

	for i := 1; i <= n; i++ {
		arr = append(arr, i*x)
	}

	return arr
}

//Another answer
func CountBy(x, n int) []int {
	var list []int
	for i := 1; i <= n; i++ {
		a := x * i
		list = append(list, a)

	}
	return list
}
