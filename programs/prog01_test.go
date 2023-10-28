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
	linkedlist.Append(300)
	linkedlist.Append(400)

	linkedlist.Traverse()
	linkedlist.Remove(400)
	linkedlist.Traverse()
}

func deleteDuplicates(head *ListNode) *ListNode {
	return nil
}

type ListNode struct {
	data int
	next *ListNode
}
type XLinkedList struct {
	head *ListNode
}

func (this *XLinkedList) Append(value int) {
	node := &ListNode{data: value}
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
