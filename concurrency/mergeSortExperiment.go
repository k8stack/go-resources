package main

import "fmt"

func main() {

	s1 := []int{3, 4, 6, 7}
	s2 := []int{10, 30, 40, 70}

	s3 := []int{1, 2, 5, 9}
	s4 := []int{20, 50, 60, 80}

	result1 := mergeSort(s1, s2)
	result2 := mergeSort(s3, s4)

	final_result := mergeSort(result1, result2)
	fmt.Print(final_result)

}

func mergeSort(s1 []int, s2 []int) []int {
	k, i, j := 0, 0, 0
	final_size := len(s1) + len(s2)
	final := make([]int, final_size, final_size)
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			final[k] = s1[i]
			i++
			k++
		} else {
			final[k] = s2[j]
			j++
			k++
		}
	}
	for i < len(s1) {
		final[k] = s1[i]
		i++
		k++

	}

	for j < len(s2) {
		final[k] = s2[j]
		j++
		k++
	}
	return final
}
