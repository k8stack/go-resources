package main

import "fmt"

type Node struct {
	val int
	left *Node
	right *Node
}

func main(){
	root := &Node{1,nil, nil}
	root.left = &Node{2, nil, nil}
	root.right = &Node{3, nil, nil}
	root.left.left = &Node{4, nil, nil}
	root.left.right = &Node{5, nil, nil}
	s := preorder(root)
	fmt.Print(s)
}

func preorder(root *Node) []int{
	var result []int
	do(root, &result)
	return result
}

func do(root *Node, s *[]int) {
        if root == nil { return }
        *s = append(*s, root.val)
	fmt.Println("result is: ", s)
	fmt.Println("result is**: ", *s)
	fmt.Println("root val is : ", root.val)
        do(root.left,s)
        do(root.right,s)
}
