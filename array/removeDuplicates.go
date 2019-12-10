package main

import (
	"fmt"
)

func main() {
	s := []int{0,0,1,1,1,1,2,2,3,3,4}
	fmt.Println("Original slice", s)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				fmt.Println("adding", s[:j], "and", s[j+1:])
				s = append(s[:j], s[j+1:]...)
				j--
				fmt.Println("after adding",s)
			}

		}
	}

	fmt.Println(s)
}

