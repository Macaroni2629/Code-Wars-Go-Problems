// https://www.codewars.com/kata/58e230e5e24dde0996000070

// Get the next prime number!

// You will get a number n (>= 0) and your task is to find the next prime number.

// Make sure to optimize your code: there will numbers tested up to about 10^12.

// Examples
// 5   =>  7
// 12  =>  13

//n = 8

func NextPrime(n int) int {
	n++ 
	for !isPrime(n)
	{n++}
	return n
  }
  

func isPrime(x int) bool { 
	if x < 4 {
	  return x == 2 || x == 3
	}
	if x % 2 == 0 { return false }
	for p := 3 ; p * p <= x ; p += 2 { 
	  if x % p == 0 { return false } 
	}
	return true
  }