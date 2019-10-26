package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		var x int
		go func() { //Block-1
			fmt.Println("####Block-1")
			x = 2
			fmt.Println("x is ",x)
		}()
		time.Sleep(100 * time.Millisecond)
		go func() { //Block-2
			fmt.Println("####Block-2")
			x = 5
			fmt.Println("x is ", x)
		}()
	}
}
