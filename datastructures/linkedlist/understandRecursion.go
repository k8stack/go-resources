package main

import "fmt"

func main() {
	recur(5)
}

func recur(x int) int {
	if x == 0 {
		return 1
	}
	fmt.Println("Before recursion ", x)
	p := recur(x - 1)
	fmt.Println("After recursion ",x)
	return p
}
