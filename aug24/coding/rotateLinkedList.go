package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func ReverseLinkedList(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func ReverseRecurr(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	res := ReverseRecurr(head.Next)
	fmt.Println(head.Val)
	head.Next.Next = head
	head.Next = nil

	return res
}
func printLinkedList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
func main() {
	head := &Node{Val: 1, Next: nil}

	curr := head
	for i := 2; i <= 5; i++ {
		curr.Next = &Node{Val: i, Next: nil}
		curr = curr.Next
	}
	printLinkedList(head)
	res := ReverseRecurr(head)
	printLinkedList(res)
}
