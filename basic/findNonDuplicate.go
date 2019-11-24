package main

import "fmt"

func main() {
	//given only one non duplicate element
	s := []int{4, 3, 2, 2, 1, 3, 1}
	var unique int
	for i := 0; i < len(s); i++ {
		unique ^= s[i]
	}
	fmt.Print(unique)
}

//https://codepumpkin.com/find-unique-array-element/#XORApproach
