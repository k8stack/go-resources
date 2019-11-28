package main

import "fmt"

func main() {

	m := make([]int, 15)
	for i := -5; i < 10; i++ {
		fmt.Println(fib(i,m))
	}

}

func fib(n int, m []int) int {

	if n <= 1 {
		return n
	}


	if m[n] > 0 {
		return m[n]
	}
	m[n] = fib(n-1, m) + fib(n-2, m)
	return m[n]
}
