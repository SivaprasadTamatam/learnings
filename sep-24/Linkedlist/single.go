package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Append(val int) {
	if ll.Head == nil {
		ll.Head = &Node{Value: val}
		return
	}

	lastNode := ll.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode.Next = &Node{Value: val}
}

func (ll *LinkedList) Delete(v int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Value == v {
		ll.Head = ll.Head.Next
		return
	}
	prev := ll.Head

	for prev.Next != nil {
		if prev.Next.Value == v {
			prev.Next = prev.Next.Next
		}
		prev = prev.Next
	}
}

func (ll *LinkedList) Print() {
	c := ll.Head

	for c != nil {
		fmt.Print(c.Value, "->")
		c = c.Next
	}
	fmt.Print("nil")
	fmt.Println()

}

func main() {
	ll := &LinkedList{}
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	ll.Print() // Output: 10 -> 20 -> 30 -> nil

	ll.Delete(20)
	ll.Print() // Output: 10 -> 30 -> nil
}
