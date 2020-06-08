# Trees

## Table of Contents

1.[Tree traversal Error with leetcode submission](#Tree-traversal-Error-with-leetcode-submission)
2.[Traversal orders](#Traversal-orders]

## Traversal orders

You can think of order as the position of the `value` with respective to `left` and `right`

- Preorder: value, left, right
- Postorder: left, right, value
- Inorder: left, value, right

## Tree traversal Error with leetcode submission

Avoid using global variables,
[read more here](https://support.leetcode.com/hc/en-us/articles/360011834174-I-encountered-Wrong-Answer-Runtime-Error-for-a-specific-test-case-When-I-test-my-code-using-this-test-case-it-produced-the-correct-output-Why-)

understand pointer well, before getting into trees & recursion, observe the below program,

```go
func main(){
    s := []int{2,2}
    var p *[]int
    p = &s
    fmt.Println(p)
    x := 2
    var y *int
    y = &x
    fmt.Println(y)
}
```

#### Output

```bash
&[2 2]
0xc0000180f0
```

### How to avoid using global variables

You can use a local variable declaration inside a function and declar another function that will call take this as parameter,
parameter should be a `pointer` to this variable, make changes

also, remember that `*&` will give value of a variable. Observe below,

```go
func main() {
    x := 2
    fmt.Println(*&x)
    fmt.Println(&x)
}
```

#### output

```go
2
0xc0000180b8
```
