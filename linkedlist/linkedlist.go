package linkedlist

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func TestLL() {
	fmt.Println("common linked list operations")
}

func (this *LinkedList) Append(data int) {
	node := &Node{Data: data}
	if this.Head == nil {
		//first node
		this.Head = node
		return
	}
	current := this.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func (this *LinkedList) Traverse() {
	current := this.Head
	for current != nil {
		fmt.Print(current.Data)
		fmt.Print(" | ")
		current = current.Next
	}
	fmt.Println()
}

func (this *LinkedList) Remove(data int) {
	if this.Head == nil {
		return
	}
	current := this.Head
	prev := current
	if this.Head.Data == data {
		this.Head = this.Head.Next //first node itself
		return
	}

	for current != nil {

		if current.Data == data {
			prev.Next = current.Next
			return

		}
		prev = current
		current = current.Next
	}

}
