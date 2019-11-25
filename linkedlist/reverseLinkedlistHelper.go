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
	head.Next.Next.Next.Next = &Node{5,nil}

	reverseLinkedlist(head)
}

func reverseLinkedlist(head *Node) *Node {
	if head.Next == nil || head == nil {
		fmt.Println("The last Element: ", head)
		return head
	}
	last := reverseLinkedlist(head.Next)
	//the statement after recursive block will execute after the recursion is complete
	fmt.Println("Head Now: ",head)
	return last
}
