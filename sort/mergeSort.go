package main

import "fmt"

func main() {
	s1 := []int{2, 4, 6, 9}
	s2 := []int{10, 13, 15, 19}

	size, i, j := len(s1)+len(s2), 0, 0

	s3 := make([]int, size, size)
	k := 0
	for i < len(s1) && j< len(s2) {
		if s1[i] < s2[j] {
			s3[k] = s1[i]
			i++
		} else {
			s3[k] = s2[j]
			j++
		}
	k++
	}

	//ignore this logic to check the difference
	for i < len(s1) {
		s3[k] = s1[i]
		i++
		k++
	}
	for j < len(s2) {
		s3[k] = s2[j]
		j++
		k++
	}

	fmt.Println(s3)

}
