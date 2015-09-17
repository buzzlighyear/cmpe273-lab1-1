


package main

import "fmt"

func fib(n uint) uint {
    if n < 2 {
        return n;
    }
    return fib(n-1) + fib(n-2);
}

func main() {
          var n uint
          fmt.Println("Enter the number whose fibonacci summation has to be printed")
          fmt.Scan(&n)
	      fmt.Println(fib(n))
		  }

