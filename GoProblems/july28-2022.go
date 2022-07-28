// https://www.codewars.com/kata/5bb904724c47249b10000131/train/go

// Our football team finished the championship. The result of each match look like "x:y". Results of all matches are recorded in the collection.

// For example: ["3:1", "2:2", "0:1", ...]

// Write a function that takes such collection and counts the points of our team in the championship. Rules for counting points for each match:

// if x > y: 3 points
// if x < y: 0 point
// if x = y: 1 point
// Notes:

// there are 10 matches in the championship
// 0 <= x <= 4
// 0 <= y <= 4

// I could not come up with an answer on my own.  Here are two answers I liked:

package kata

import ( "strings" )

func Points(games []string) int {
  result := 0
  for _, game := range games {
    str := strings.Split(game, ":")
    x, y := str[0], str[1]
    switch {
      case x > y:
        result += 3
      case x == y:
        result += 1
    }
  }
  return result
}

// This first answer was top voted "best practices" because it uses the `strings.Split` method.  It splits on the colon.
//Then it initializes and declares `x` and `y` variables.

//Here is a second answer

package kata

func Points(games []string) int {
  points := 0
  for _, g := range games {
      if g[0]>g[2] {
        points += 3
      } else if g[0] == g[2] {
        points += 1
      }
  }
  return points
}

//As this iterates through the slice of strings, it compares g[0] to g[2].  Luckily, in Go, you can compare string versions of numbers.