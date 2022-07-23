// https://www.codewars.com/kata/5ce9c1000bab0b001134f5af/train/go

// Given a month as an integer from 1 to 12, return to which quarter of the year it belongs as an integer number.

// For example: month 2 (February), is part of the first quarter; month 6 (June), is part of the second quarter; and month 11 (November), is part of the fourth quarter.

//Answer that I wanted but couldn't come up with completely due to lack of mastery of switch syntax

package kata

func QuarterOf(month int) (result int) {
  
  switch month{
    case 1,2,3:result = 1
    case 4,5,6:result = 2
    case 7,8,9:result = 3
    case 10,11,12:result = 4
  }
  return
}

// a clever one

package kata

func QuarterOf(month int) int {
  return (month + 2) / 3
}

//this answer takes advantage of that when you do divison, remainders are completely ignored