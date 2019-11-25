package main

import "fmt"

type Node struct {
	Val int
	Next *Node
}
func main(){
	head := &Node{1,nil}
	head.Next = &Node{2,nil}
	head.Next.Next = &Node{3,nil}
	head.Next.Next.Next = &Node{4,nil}

	result := reverseLinkedlist(head)
	fmt.Println("result is")
	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
}

func reverseLinkedlist(head *Node) *Node{
	fmt.Println("=======",*head)
	if head.Next == nil || head == nil {
		return head
	}
	p := reverseLinkedlist(head.Next)
	fmt.Println("*******",*head)
	fmt.Println("head is",*head)
	fmt.Println("P is",*p)
	fmt.Println("before", head.Next.Next)
	head.Next.Next = head
	fmt.Println("after", head.Next.Next)
	head.Next = nil
	return p
}
