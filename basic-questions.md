## Golang interview questions
### Table of Contents
1. [What are different data types in go](#what-are-different-data-types-in-go)
2. [what is a package in go](#what is a package in go)
3. [what is go workspace](#what is go workspace)
4. [what is GOPATH](#what is GOPATH)
5. [what are different directories inside a go project](#what are different directories inside a go project)
6. [What are different data structures in go](#What are different data structures in go)
7. [How to iterate maps in go](#How to iterate maps in go)
12. [How do you create go modules](#How do you create go modules)
13. [Explain go modules](#Explain go modules)
14. [Function syntax in go](#Function syntax in go)
15. [Method syntax in go](#Method syntax in go)
16. [How do you declare inline function](#How do you declare inline function)
17. [What is the use of empty interface](#What is the use of empty interface)
18. [How access modifiers work in go](#How access modifiers work in go)
19. [what is GOPATH and GOROOT](#what is GOPATH and GOROOT)
20. [Explain polymorphism in go](#Explain polymorphism in go)
21. [Explain different print formats](#Explain different print formats)
22. [What is defautl value of a global, local, & pointer variable](#What is defautl value of a global, local, & pointer )variable
23. [Does go support method overloading, operator overloading, type inheritance](#Does go support method overloading, operator )overloading, type inheritance
24. [Write a program on pointers](#Write a program on pointers)
25. [How do you copy slices](#How do you copy slices)
26. [How do you copy interfaces](#How do you copy interfaces)
27. [How do you compare two structs](#How do you compare two structs)
28. [How do you compare two interfaces](#How do you compare two interfaces)
29. [Explain go get command](#Explain go get command)
30. [How do you do indentation](#How do you do indentation)

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

##### Explain go get command
go get is used for installations,
suppose if you are installing golint, you do,
```go
go get -u golang.org/x/lint/golint
```
`the source code of `golint` will be downloaded and copied to $GOPATH/src/golang.org/x/lint`
and the binary `golint` will be copied to `$GOPATH/bin`
since this is added to PATH	during installation, the binary will be
available to the system

##### How do you do indentation
using `go fmt`,
example,
```go
go fmt hello.go
```
go fmt by default uses tabs & spaces are not longer used/recommended.

##### How do you prevent a race condtion?
