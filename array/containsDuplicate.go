package main

import "fmt"

func main() {

	s := []int{1, 2, 1, 3}
	result := containsDuplicate(s)
	fmt.Println(result)
}

func containsDuplicate(s []int) bool {
	m := make(map[int]bool)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok == true {
			return true
		} else {
			m[s[i]] = true
		}
	}
	return nil
}
