# Golang Concurrency

## Table of Contents
1. [What is thread safety?](#what-is-thread-safety)
2. [What is channel](#what-is-channel)
3. [What is buffered channel?](#what-is-buffered-channel)
4. [Channel synchronization](#channel-synchronization)
5. [What is the use of Mutex?](#what-is-the-use-of-mutex)
6. [Explain channel blocking](#explain-channel-blocking)
7. [How do you make a function thread safe in golang](#How do you make a function thread safe in golang)
8. [Explain deadlock condition](#what-is-a-package-in-go)
9. [What is the difference between concurrency and parallelism?](#what-is-the-difference-between-concurrency-and-parallelism)

## What is thread safety?

## What is channel?
A Channel is a  pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.

```go
package main

import "fmt"

func main() {

    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)
}
```

## What is buffered channel?

![](https://image.slidesharecdn.com/golang101-170518032637/95/golang-101-17-638.jpg?cb=1495078019)

```go
package main

import "fmt"

func main() {

    messages := make(chan string, 2)

    messages <- "1"
    messages <- "2"
    messages <- "3"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
```

## Channel synchronization


```go
package main

import "fmt"
import "time"

func task1(done chan bool) {
    fmt.Print("Task 1")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func task2() {
    fmt.Println("Task 2 (goroutine)")
}

func main() {
   done := make(chan bool,1)
   go task1(done)


   if <- done {
       go task2()
       fmt.Scanln()
   }
}
```

## What is the use of Mutex?
A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point of time to prevent race condition from happening

```go
mutex.Lock()  
x = x + 1  
mutex.Unlock()  
```

__With Race Condition__
```go
package main  
import (  
    "fmt"
    "sync"
    )
var x  = 0  
func increment(wg *sync.WaitGroup) {  
    x = x + 1
    wg.Done()
}
func main() {  
    var w sync.WaitGroup
    for i := 0; i < 1000; i++ {
        w.Add(1)        
        go increment(&w)
    }
    w.Wait()
    fmt.Println("final value of x", x)
}
```

## Explain channel blocking 



## How do you make a function thread safe in golang?
Using Mutex

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

var lock sync.Mutex

func main() {
    go importantFunction("first")
    go importantFunction("second")
    time.Sleep(3 * time.Second)
}


func importantFunction(name string) {
    lock.Lock()
    defer lock.Unlock()
    fmt.Println(name)
    time.Sleep(1 * time.Second)
}
```
  
## Explain deadlock condition

If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. 
If this does not happen, then the program will panic at runtime with Deadlock.

```go
package main


func main() {  
    ch := make(chan int)
    ch <- 5
}
```

## What is the difference between concurrency and parallelism?

![](http://www.yosefk.com/img/n/concurrency-centric.png)

## Unidirectional channels

```go
package main

import "fmt"

func sendData(sendch chan<- int) {  
    sendch <- 10
}

func main() {  
    sendch := make(chan<- int)
    go sendData(sendch)
    fmt.Println(<-sendch)
}
```
