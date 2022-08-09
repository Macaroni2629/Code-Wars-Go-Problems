// https://www.codewars.com/kata/57a5c31ce298a7e6b7000334/train/go

// Complete the function which converts a binary number (given as a string) to a decimal number.

func BinToDec(bin string) int {
	n := 0
	for _, r := range bin {
	  n *= 2
	  n += int(r-'0')
	}
	return n
  }

  //another answer

  func BinToDec(bin string) int {
	r, _ := strconv.ParseInt(bin, 2, 64)
	return int(r)
 }