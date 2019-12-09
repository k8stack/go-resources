# Golang Concurrency

## Table of Contents
1. [What is thread safety?](#what-is-thread-safety)
2. [What is channel](#what-is-channel)
3. [What is buffered channel?](#what-is-buffered-channel)
4. [Channel synchronization](#channel-synchronization)
5. [What is the use of Mutex?](#what-is-the-use-of-mutex)
6. [Explain channel blocking](#explain-channel-blocking)
7. [How do you make a function thread safe in golang](#How-do-you-make-a-function-thread-safe-in-golang)
8. [Explain deadlock condition](#what-is-a-package-in-go)
9. [What is the difference between concurrency and parallelism?](#what-is-the-difference-between-concurrency-and-parallelism)
8. [Closing channels](#closing-channels)

## What is thread safety?
Thread-safe code only manipulates shared data structures in a manner that ensures that all threads behave properly and fulfill their design specifications without unintended interaction.

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
####Buffered vs Unbuffered
![](https://www.ardanlabs.com/images/goinggo/86_guarantee_of_delivery.png)

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

**Read Blocking**
If you are trying to read data from a channel but channel does not have a value available with it, it blocks the current goroutine and unblocks other in a hope that some goroutine will push a value to the channel. Hence, this read operation will be blocking

**Send Blocking**
Similarly, if you are to send data to a channel, it will block current goroutine and unblock others until some goroutine reads the data from it. Hence, this send operation will be blocking

**Default behaviour**
Sends and receives to a channel are blocking by default. What does this mean? When a data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel. Similarly when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.

This property of channels is what helps Goroutines communicate effectively without the use of explicit locks or conditional variables that are quite common in other programming languages

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
![](https://miro.medium.com/max/601/1*esw_FDWZXB-3o2gA4buRpA.jpeg)
![](https://miro.medium.com/max/601/1*cZY2BoVrw7OSpl2-r-g5yA.jpeg)

## What is the difference between concurrency and parallelism?

![](http://www.yosefk.com/img/n/concurrency-centric.png)

## Unidirectional channels

By default, channels are bi-directional (send-recieve). They can be made uni-directional (send-only/recieve-only)

```
ch1 := make(<-chan int) //read-only
ch2 := make(chan<- int) //send-only
```

But what is the use of unidirectional channel? Using unidirectional channels increases the type-safety of a program. Hence the program is less prone to error.


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

## Closing channels

A channel can be closed so that no more data can be sent through it.

How to check if a channel is closed or open?
```
val, ok := <- channel
```
ok -> true //channel is open
ok -> false //chananel is closed

```go
package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
}
```
