package programs

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic ", r)
		}
	}()
	tests()
}

func tests() {
	linkedlist := &XLinkedList{}
	linkedlist.Append(100)
	linkedlist.Append(200)
	linkedlist.Append(200)
	linkedlist.Append(300)
	linkedlist.Append(400)

	linkedlist.Traverse()

	linkedlist.Remove(400)

}
// delete duplicate from a sorted list
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

type Node struct {
	data int
	next *Node
}
type ListNode struct {
	Val  int
	Next *ListNode
}
type XLinkedList struct {
	head *Node
}

func (this *XLinkedList) Append(value int) {
	node := &Node{data: value}
	if this.head == nil {
		this.head = node
		return
	}
	current := this.head
	for current.next != nil {
		current = current.next
	}
	current.next = node

}
func (this *XLinkedList) Traverse() {
	fmt.Println()
	current := this.head
	for current != nil {
		fmt.Print(current.data)
		fmt.Print(" | ")
		current = current.next
	}
}
func (this *XLinkedList) Remove(val int) {
	if this.head == nil {
		return
	}
	if this.head.data == val {
		this.head = this.head.next
		return
	}
	current := this.head
	prev := this.head
	for current != nil {
		if current.data == val {
			current = current.next
			prev.next = current
			return
		}
		prev = current
		current = current.next
	}

}
