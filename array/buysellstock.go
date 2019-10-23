package main

import "fmt"

func main() {
	s := []int{7,1,5,3,6,4}
	profit := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i+1] - s[i] > 0 {
			profit+= s[i+1] - s[i]
		}
	}
	fmt.Println(profit)
}
