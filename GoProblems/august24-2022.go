// https://www.codewars.com/kata/57a6633153ba33189e000074

// Count the number of occurrences of each character and return it as a list of tuples in order of appearance. For empty output return an empty list.

// Example:

// OrderedCount("abracadabra") == []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}

// // Where
// type Tuple struct {
//     Char  rune
//     Count int
// }

package orderedcount

// Use the preloaded Tuple struct as return type
// type Tuple struct {
//  Char  rune
//  Count int
// }

import (
	"strings"
)

func OrderedCount(text string) []Tuple {
	// Implement me! :)
	var TupleList []Tuple = []Tuple{}
	var hashMap = make(map[rune]int)
	for _, c := range text {
		if _, ok := hashMap[c]; !ok {
			v := strings.Count(text, string(c))
			hashMap[c] = v
			TupleList = append(TupleList, Tuple{Char: c, Count: v})
		}
	}

	return TupleList
}
