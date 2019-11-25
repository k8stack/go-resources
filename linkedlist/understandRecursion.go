package main

import "fmt"

func main() {
	fmt.Print(recur(5))
}

func recur(x int) int {
	if x == 0 {
		return 1
	}
	fmt.Println(x)
	p := recur(x - 1)
	return p
}
