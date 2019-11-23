package main

import "fmt"

func main() {
	s1 := []int{2, 4, 6, 9}
	s2 := []int{1, 3, 5, 11}

	size, i, j := len(s1)+len(s2), 0, 0

	s3 := make([]int, size, size)
	k := 0
	for k = 0; k < size-1; k++ {
		if s1[i] < s2[j] {
			s3[k] = s1[i]
			i++
		} else {
			s3[k] = s2[j]
			j++
		}

	}

	//ignore this logic to check the difference
	for i < len(s1) {
		s3[k] = s1[i]
		i++
	}
	for j < len(s2) {
		s3[k] = s2[j]
		j++
	}

	fmt.Println(s3)

}
