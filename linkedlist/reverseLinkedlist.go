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
	//note: below condition failed on leetcode submission, 
	//if head.Next == nil || head == nil {
	//error: Line 9: panic: runtime error: invalid memory address or nil pointer dereference

	if head == nil || head.Next == nil {
		return head
	}
	p := reverseLinkedlist(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
