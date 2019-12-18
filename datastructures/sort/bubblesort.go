package main

import "fmt"

func main() {

	s := []int{6, 9, 8, 1, 2, 3, 7}
	fmt.Println("Before: ", s)
	bubblesort(s)
	fmt.Println("After: ", s, "\n")
}

func bubblesort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
