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



