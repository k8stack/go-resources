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


/* explanation

1 -> 2 -> 3 -> 4 -> 5 -> x
=====================================================================
1 -> 2 -> 3 -> 4 -> 5 -> 4 //head is at 4 & head.next.next = head
1 -> 2 -> 3 -> 4 -> x //head.next = nil
5 -> 4 -> x
=====================================================================
1 -> 2 -> 3 -> 4 -> 3 //head is at 3 & head.next.next = head
1 -> 2 -> 3 -> x //head.next = nil
5 -> 4 -> 3 -> x
=====================================================================
1 -> 2 -> 3 -> 2 //head is at 2 & head.next.next = head
1 -> 2 -> x //head.next = nil
5 -> 4 -> 3 -> 2 -> x
=====================================================================
1 -> 2 -> 1 //head is at 1 & head.next.next = head
1 -> x //head.next = nil
5 -> 4 -> 3 -> 2 -> 1 -> x

/*
