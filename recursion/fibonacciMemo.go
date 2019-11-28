package main

import "fmt"

func main() {

	m := make([]int, 30)
	for i := 0; i < 30; i++ {
		fmt.Println(fib(i,m))
	}

}

func fib(n int, m []int) int {

	if n < 0 {
		return n
	}

	if n == 0 || n == 1 {
		return 1
	}

	if m[n] > 0 {
		return m[n]
	}
	m[n] = fib(n-1, m) + fib(n-2, m)
	return m[n]
}
