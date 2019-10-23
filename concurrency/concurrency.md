## Golang Concurrency questions
### Table of Contents

1. [What is a go routine](#What is a go routine)
2. [How to create go routine in go](#How to create go routine in go)
3. [Race condition program](#Race condition program)
4. [How do you prevent a race condtion](#How do you prevent a race condtion)
5. [What are channels in go](#What are channels in go)

##### What is a go routine?
A goroutine is a  lightweight thread managed by the Go runtime

##### How to create go routine in go?
using `go` keyword

```go
package main

import (
	"fmt"
)

func main() {
		go fmt.Println("Hello")
}

```

##### Race condition program
```go
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

```
##### How do you prevent a race condtion?

##### What are channels in go?
A channel is a data transfer pipe between go routines where
data can be passed into and read from.

*a channel can transport data of only one data type.*

##### Channel example
```go
func main(){
	c:=make(chan int)
    fmt.Println("%T", c)
    fmt.Println("%V", c)
}
```
output will be a memory address,
*Channels by default are pointers*

##### What are bufferred and unbuffered channels?
```go
// unbuffered channel of ints
ic := make(chan int)

// buffered channel with room for 10 strings
sc := make(chan string, 10)
```

##### How do you close channels?
using 	`close(ch)`

