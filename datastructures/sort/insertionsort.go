package main

import "fmt"

func main() {

	s := []int{8, 2, 4, 1, 9, 3, 6}
	fmt.Println("Before: ", s)
	insertionsort(s)
	fmt.Println("After: ", s, "\n")
}

func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
