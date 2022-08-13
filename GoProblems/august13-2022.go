// https://www.codewars.com/kata/57a1d5ef7cb1f3db590002af/train/go

// Create function fib that returns n'th element of Fibonacci sequence (classic programming task).

func Fib(n int) int {
	a,b := 0,1
	for ;n>0; n-- {
	  a,b = b,a+b
	}
	return a
  }


func Fib(n int) int {
	a := 0
	b := 1
	for ; n > 0; n-- {
		a = b
		b = a + b
	}
	return a
}

func Fib(n int) int {
	if n == 0{
	  return 0
	}
	if n <= 2{
	  return 1
	}
	  return Fib(n-1) + Fib(n-2)
	}