## Golang interview questions
### Table of Contents
1. [What are different data types in go?](#what-are-different-data-types-in-go)
2. [How do you copy maps](#how-do-you-copy-maps)

## This is the introduction <a name="introduction"></a>
Some introduction text, formatted in heading 2 style

## Some paragraph <a name="paragraph1"></a>
The first paragraph text

### Sub paragraph <a name="subparagraph1"></a>
This is a sub paragraph, formatted in heading 3 style

## Another paragraph <a name="paragraph2"></a>
The second paragraph text

#### What are different data types in go?
int8(aka byte), int16, int32(aka rune), int64
float64,
byte,
bool,
string

##### what is a package in go?
a package is a directory where all go files resides. 

##### what is go workspace?
a workspace is a directory heirarchy with two root directories,
 - src -> go source files (typically contains multiple vcs based directories example, github.com)
 - bin -> executable files
 - pkg -> shared libs used by executables, example: go mod dependencies 
##### what is GOPATH?
It specifies the go workspace 

##### what are different directories inside a go project?
- cmd -> will have sub-directories for cli based
- pkg -> if you need your code to be re-used by other projects (careful with this)
- internal -> if you want your code to be private 
- api -> openapi/swagger
- web -> web components
- configs -> configurations
- init -> system init/process manager (systemd, sysv etc..)
- scripts -> build/installs
- build -> for deb,rpm,ami,docker images etc..
- deployments -> docker-compose, k8s yaml, terraform etc. 
- test -> testing 

##### What are different data structures in go?
Array - fixed length
Slice - variable length
Maps - key values
Struct - is a collection of fields of same/diff types

##### How to iterate maps in go?

using range keyword
```go
  m := map[string]string{ "key1":"val1", "key2":"val2" };
  for k, v := range m {
    fmt.Printf("key[%s] value[%s]\n", k, v)
  }
```

##### What is a go routine ?
A goroutine is a  lightweight thread managed by the Go runtime


##### How to create go routine in go?

##### What are channels in go?

##### Explain race condition in go?

##### How do you create go modules?
using init command
```go
go module init
````

##### Explain go modules
Go module is a dependency management system introduced from go v1.11
a module is a collection of go packages stored in `go.mod` file

##### Function syntax in go?
```go
func add(x int32, y int32) return int32 {
  return x+y
}
```
##### Method syntax in go?

```go
package main

import "fmt"

type Cal struct {
	x int32
	y int32
}

func (*Cal) add(x int32, y int32) int32 {
	return x + y
}

func main() {
	c := &Cal{}
	c.add(2, 2)
	fmt.Println(c.add(2, 2))
}
```

##### How do you declare inline function?
inline function is an anonymous function used 

Simple 
```go
func(){ 

  fmt.Println("Hello") 
}() 
	
```
Assigning
```go
value := func(){ 
  fmt.Println("Welcome! to GeeksforGeeks") 
} 
value() 
```

##### What is the use of empty interface?
Interface has two meanings 
 - interface is a set of method declaration
 - emtpy interface is a type used to dynamic type conversion
   - Example, 
     You can pass int or float, its dynamic
     ```
     func add (x interface{}, y interface{}) {
       x+y
     }
     ```

##### How access modifiers work in go?
there are two types of access modifiers
  - exported (Accessible outside package)
  - unexported (Accessible only within the package)

USING UPPERCASE WILL EXPORT A METHOD OR VARIABLE TO OUTSIDE PACKAGE

```go
  Foo int; // exported - accessible outside package
  bar int; // unexported - only accessible wit
```

##### what is GOPATH and GOROOT?

GOPATH specifies go workspace and GOROOT specifies go installation directory

##### Explain polymorphism in go?

Polymorphism can be achieved using interface

```go
package main

import "fmt"

type User struct {
	username string
	email string
}

type UserI interface {
	getEmail() string
}

func (u *User) getEmail() string {
	return u.username
}

func main(){
	var userI UserI
	userI = &User{"zillani", "shaikzillani@gmail.com"}
	fmt.Println(userI.getEmail())
}
```

##### Explain different print formats

##### What is defautl value of a global, local, & pointer variable?
  global -> 0
  local -> 0
  pointer -> nil

##### Does go support method overloading, operator overloading, type inheritance?

##### Write a program on pointers

Swapping two variables 

```go
package main

import "fmt"

func swap(px, py *int){
  temp := *px
  *px = *py
  *py = temp
}

func main(){
  a:=1
  b:=2

  fmt.Println("Before swapping: ",a,b)
  swap(&a,&b)
  fmt.Println("After swapping: ",a,b)
}

```

#### How do you create slice?

using MAKE or using array declaration
```go
s:= make([]int,3)
s[0] = 1
s[1] = 2
s[2] = 3

s2 := []int{1,2,3}
fmt.Println(s,s2)
```
##### How do you copy slices?
```go
s1 := []int{1, 2, 3}
s2 := make([]int, 3)
copy(s2, s1)
fmt.Println(s1)
fmt.Println(s2)
```
## How do you copy maps?

##### How do you copy interfaces?

##### How do you compare two structs?

##### How do you compare two interfaces?
