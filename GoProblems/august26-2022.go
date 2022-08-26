// https://www.codewars.com/kata/56b97b776ffcea598a0006f2

// Task
// Given an array of integers, your function bubblesortOnce/bubblesort_once/BubblesortOnce (or equivalent, depending on your language's naming conventions) should return a new array equivalent to performing exactly 1 complete pass on the original array. Your function should be pure, i.e. it should not mutate the input array.

func BubblesortOnce(s []int) []int {
	numbers := make([]int, len(s))
	copy(numbers, s)
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
		}
	}
	return numbers
}
  