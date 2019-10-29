package main

import "fmt"

func main() {
	for i:=1; i<10; i++ {
		fmt.Println(fib(i))
	}
}

func fib(n int) int {
	if n >= 3 {
		return fib(n-1) + fib(n-2)
	} else {
		return 1
	}
}
