## Golang interview questions
### Table of Contents
1. [What are different data types in go](#what-are-different-data-types-in-go)
2. [what is a package in go](#what-is-a-package-in-go)
3. [what is go workspace](#what-is-go-workspace)
4. [what is GOPATH](#what-is-GOPATH)
5. [what are different directories inside a go project](#what-are-different-directories-inside-a-go-project)
6. [What are different data structures in go](#What-are-different-data-structures-in-go)
7. [How to iterate maps in go](#How-to-iterate-maps-in-go)
8. [How do you create go modules](#How-do-you-create-go-modules)
9. [Explain go modules](#Explain-go-modules)
10. [Function syntax in go](#Function-syntax-in-go)
11. [Method syntax in go](#Method-syntax-in-go)
12. [How do you declare inline function](#How-do-you-declare-inline-function)
13. [What is the use of empty interface](#What-is-the-use-of-empty-interface)
14. [How access modifiers work in go](#How-access-modifiers-work-in-go)
15. [what is GOPATH and GOROOT](#what-is-GOPATH-and-GOROOT)
16. [Explain polymorphism in go](#Explain-polymorphism-in-go)
17. [Explain different print formats](#Explain-different-print-formats)
18. [What is defautl value of a global, local, & pointer variable](#What-is-defautl-value-of-a-global,-local,-&-pointer-variable)
19. [Does go support method overloading, operator overloading, type inheritance](#Does-go-support-method-overloading,-operator-overloading,-type-inheritance)
20. [Write a program on pointers](#Write-a-program-on-pointers)
21. [How do you copy slices](#How-do-you-copy-slices)
22. [How do you copy interfaces](#How-do-you-copy-interfaces)
23. [How do you compare two structs](#How-do-you-compare-two-structs)
24. [How do you compare two interfaces](#How-do-you-compare-two-interfaces)
25. [Explain go get command](#Explain-go-get-command)
26. [How do you do indentation](#How-do-you-do-indentation)
27. [Explain closures in go](#Explain-closures-in-go)
28. [How to increase slice capacity](#how-to-increase-slice-capacity?)
29. [How to convert string to int](#How-to-convert-string-to-int)
30. [What is & and * in go](#What-is-&-and-*-in-go)
31. [Explain ++ and -- in golang](#Explain-++-and----in-golang)
32. [Does go support indexing](#Does-go-support-indexing)
33. [How to convert string to byte array & viceversa](#How-to-convert-string-to-byte-array-&-viceversa)

#### What are different data types in go?
int8(aka byte), int16, int32(aka rune), int64
float64,
byte,
bool,
string

#### what is a package in go?
a package is a directory where all go files resides.

#### what is go workspace?
a workspace is a directory heirarchy with two root directories,
 - src -> go source files (typically contains multiple vcs based directories example, github.com)
 - bin -> executable files
 - pkg -> shared libs used by executables, example: go mod dependencies

#### what is GOPATH?
It specifies the go workspace

#### what are different directories inside a go project?
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

#### What are different data structures in go?
Array - fixed length
Slice - variable length
Maps - key values
Struct - is a collection of fields of same/diff types

#### How to iterate maps in go?

using range keyword
```go
  m := map[string]string{ "key1":"val1", "key2":"val2" };
  for k, v := range m {
    fmt.Printf("key[%s] value[%s]\n", k, v)
  }
```

#### How do you create go modules?
using init command
```go
go module init
````

#### Explain go modules
Go module is a dependency management system introduced from go v1.11
a module is a collection of go packages stored in `go.mod` file

#### Function syntax in go?
```go
func add(x int32, y int32) return int32 {
  return x+y
}
```
#### Method syntax in go?

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

#### How do you declare inline function?
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

#### What is the use of empty interface?
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

#### How access modifiers work in go?
there are two types of access modifiers
  - exported (Accessible outside package)
  - unexported (Accessible only within the package)

USING UPPERCASE WILL EXPORT A METHOD OR VARIABLE TO OUTSIDE PACKAGE

```go
  Foo int; // exported - accessible outside package
  bar int; // unexported - only accessible wit
```

#### what is GOPATH and GOROOT?

GOPATH specifies go workspace and GOROOT specifies go installation directory

#### Explain polymorphism in go?

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

#### Explain different print formats

#### What is defautl value of a global, local, & pointer variable?
  global -> 0
  local -> 0
  pointer -> nil

#### Does go support method overloading, operator overloading, type inheritance?

#### Write a program on pointers

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
#### How do you copy slices?
```go
s1 := []int{1, 2, 3}
s2 := make([]int, 3)
copy(s2, s1)
fmt.Println(s1)
fmt.Println(s2)
```
## How do you copy maps?

#### How do you copy interfaces?

#### How do you compare two structs?

#### How do you compare two interfaces?

#### Explain go get command
go get is used for installations,
suppose if you are installing golint, you do,
```go
go get -u golang.org/x/lint/golint
```
`the source code of `golint` will be downloaded and copied to $GOPATH/src/golang.org/x/lint`
and the binary `golint` will be copied to `$GOPATH/bin`
since this is added to PATH	during installation, the binary will be
available to the system

#### How do you do indentation
using `go fmt`,
example,
```go
go fmt hello.go
```
go fmt by default uses tabs & spaces are not longer used/recommended.

#### Explain closures in go

```go
package main
import "fmt"
var global func()
func closure() {
 var A int = 1
 func() {
  var B int = 2
  func() {
   var C int = 3
   global = func() {
    fmt.Println(A, B, C)
    fmt.Println(D, E, F) // causes compilation error
   }
   var D int = 4
  }()
  var E int = 5
 }()
 var F int = 6
}
func main() {
 closure()
 global()
}
```
**global function has access to,**
- It has access to A,B,C since the layer of enclosure does not matter
- It doesn’t have access to D,E,F since their declaration don’t precede
- Even after execution of “closure” function, it’s local variables were not destroyed. They were accessible in the call to “global” function


#### How to increase slice capacity?
```go
s = s[:cap(s)]
```

#### How to convert string to int?

```go
package main

import (
	"fmt"
	"strconv"
)
s := "123"
num,_ := strconv.Atoi(s)
fmt.Println(num)
```

#### What is & and * in go?

__The & Operator__

`&` gives address of a variable

__The * Operator__

\\`*` used to hold the address of a variable

```go
func main(){
        var x int = 2;
        var y *int = &x
        fmt.Println("Address of x ",y)
}
```
![](https://media.geeksforgeeks.org/wp-content/uploads/20190705160332/Pointers-in-Golang.jpg)

#### Explain ++ and -- in golang?

++ and -- are statements but not expressions
[check here](https://stackoverflow.com/questions/25800242/go-golang-syntax-error-unexpected-expecting)


#### Does go support indexing?

[check here](https://flaviocopes.com/golang-does-not-support-indexing/)

#### How to convert string to byte array & viceversa?

__string to byte array__
```go
b := []byte("ABC€")
fmt.Println(b) // [65 66 67 226 130 172]
```

__byte array to string__
```go
s := string([]byte{65, 66, 67, 226, 130, 172})
fmt.Println(s)
```
