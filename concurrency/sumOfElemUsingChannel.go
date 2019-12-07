package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := s[:len(s)/2]
	s2 := s[len(s)/2:]

	fmt.Println(s1, s2)

	c := make(chan int)

	go add(0, s1, c)

	result := <-c

	go add(result, s2, c)

	result = <-c

	fmt.Println(result)
}

func add(i int, s []int, c chan int) {
	var result int
	for i := 0; i < len(s); i++ {
		result += s[i]
	}
	c <- i + result
}
