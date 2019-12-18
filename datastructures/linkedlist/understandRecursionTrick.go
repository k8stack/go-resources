package main

import "fmt"

func main(){
	do(1)
}

//understand how this traverses backwards, the same logic is applied for printing the linkedlist in reverse order, 
//only after the recrusion is complete, the statements after it are executed.
func do(n int) int {
	fmt.Println("loop")
	if n == 5 {
		return 5
	}
	do(n+1)
	fmt.Println(n)
	return n
}
