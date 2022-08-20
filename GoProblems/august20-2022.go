// https://www.codewars.com/kata/57f780909f7e8e3183000078

// Given a non-empty array of integers, return the result of multiplying the values together in order. Example:

// [1, 2, 3, 4] => 1 * 2 * 3 * 4 = 24

func Grow(arr []int) int {
	result := 1

	for _, n := range arr {
		result *= n
	}

	return result
}

func Grow(arr []int) int {
	v := arr[0]

	for _, val := range arr[1:] {
		val *= val
	}

	return v
}