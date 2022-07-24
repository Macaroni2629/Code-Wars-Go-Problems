//https://www.codewars.com/kata/5672a98bdbdd995fad00000f/train/go

// Rock Paper Scissors
// Let's play! You have to return which player won! In case of a draw return Draw!.

// Examples(Input1, Input2 --> Output):

// "scissors", "paper" --> "Player 1 won!"
// "scissors", "rock" --> "Player 2 won!"
// "paper", "paper" --> "Draw!"

//My solution 

package kata


func Rps(p1, p2 string) string{
  var answer string = ""
  switch {
  case p1 == "scissors" && p2 == "paper":
    answer = "Player 1 won!"
  case p1 == "scissors" && p2 == "rock":
    answer = "Player 2 won!"
  case p1 == "rock" && p2 == "scissors":
    answer = "Player 1 won!"
  case p1 == "rock" && p2 == "paper":
    answer = "Player 2 won!"
  case p1 == "paper" && p2 == "rock":
    answer = "Player 1 won!"
  case p1 == "paper" && p2 == "scissors":
    answer = "Player 2 won!"
  default:
    answer = "Draw!"
  }
  return answer
}

//Solution I like

package kata

var m = map[string]string{"rock": "paper", "paper": "scissors", "scissors": "rock"}

func Rps(a, b string) string {
  if a == b {
    return "Draw!"
  }
  if m[a] == b {
    return "Player 2 won!"
  }
  return "Player 1 won!"
}

// I like that this person used a `map` data structure to abstract away the repetitiveness of the code.