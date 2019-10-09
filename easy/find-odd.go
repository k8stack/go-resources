package main

import "fmt"

func main() {
a := []int{1, 2, 3, 4, 5, 6}
  for num := range a {
    if num%2 != 0 {
	fmt.Print(num)
    }
  }
}

