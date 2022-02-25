package main

import "fmt"

func main() {
	s := []int{2, 7, 11, 15}
	target := 9
	result := twoSums(s, target)
	fmt.Println(result)
}

func twoSums(s []int, target int) []int {
	// put array into a map like,
	// | 2, 0 | 7, 1 | 11, 3 | 15, 4 |
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}
	// substract target from array each value 
	// if the result exists in the map then return
	for i := 0; i < len(s); i++ {
		val, isExist := m[target-s[i]]

		if isExist && val != i {
			return []int{i, m[target-s[i]]}
		}
	}
	return nil

}
