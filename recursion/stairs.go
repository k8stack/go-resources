package main

import "fmt"

func main() {
	n := 5
	m := make([]int, n+1, n+1)
	result := climbStairs(n,m)
	fmt.Println(result)
}
func climbStairs(n int, m []int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if m[n] > 0 {
		return m[n]
	}
	m[n] = climbStairs(n-1, m) + climbStairs(n-2, m)
	return m[n]
}
