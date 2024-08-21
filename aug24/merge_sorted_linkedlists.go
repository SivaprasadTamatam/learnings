package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func mergeLists(l1, l2 *Node) *Node {
	res := &Node{}

	curr := res

	for l1 != nil && l2 != nil {
		if l1.val < l2.val {
			curr.next = l1
			l1 = l1.next
		} else {
			curr.next = l2
			l2 = l2.next
		}
		curr = curr.next
	}

	if l1 != nil {
		curr.next = l1
	}

	if l2 != nil {
		curr.next = l2
	}

	return res.next
}

func Print(n *Node) {
	for n != nil {
		fmt.Println(n.val)
		n = n.next
	}
}

func main() {
	l1 := &Node{1, &Node{2, &Node{5, nil}}}
	l2 := &Node{3, &Node{4, &Node{6, nil}}}

	Print(mergeLists(l1, l2))
}
