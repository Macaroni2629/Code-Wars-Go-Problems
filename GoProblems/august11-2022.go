// https://www.wsj.com/articles/walgreens-pays-bonuses-up-to-75-000-to-recruit-pharmacists-11660227605?reflink=desktopwebshare_permalink

// Return the Nth Even Number

// Example(Input --> Output)

// 1 --> 0 (the first even number is 0)
// 3 --> 4 (the 3rd even number is 4 (0, 2, 4))
// 100 --> 198
// 1298734 --> 2597466
// The input will not be 0.

func NthEven(n int) int {
    return 2*(n-1)
}

// https://www.codewars.com/kata/57cfdf34902f6ba3d300001e/train/go

// You will be given a list of strings. You must sort it alphabetically (case-sensitive, and based on the ASCII values of the chars) and then return the first value.

// The returned value must be a string, and have "***" between each of its letters.

// You should not remove or add elements from/to the array.

import (
	"sort"
	"strings"
  )
  
  func TwoSort(arr []string) string {
	sort.Strings(arr)
	chars := strings.Split(arr[0], "")
	return strings.Join(chars, "***")
  }