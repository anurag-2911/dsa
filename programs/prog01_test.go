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
func merge02(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	k := n - 1
	diff := len(nums1) - m
	ln := len(nums1) - 1

	for counter := 0; counter < diff; {
		if nums1[i] > nums2[k] {
			nums1[ln] = nums1[i]
			ln--
			i--
		} else {
			nums1[ln] = nums2[k]
			k--
			ln--
		}
	}
	//remaining
	for ; i > 0; i-- {
		nums1[ln] = nums1[i]
		ln--
	}
	for ; k > 0; k-- {
		nums1[ln] = nums2[k]
		ln--
	}
}
func TestMergeArrays(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	// merge(nums1, m, nums2, n)
	merge02(nums1, m, nums2, n)

	nums11 := []int{1}
	m1 := 1
	nums22 := []int{}
	n1 := 0
	// merge(nums11,m1,nums22,n1)
	merge02(nums11, m1, nums22, n1)
}

/*
*You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function,
but instead be stored inside the array nums1. To accommodate this,
nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	resultarray := make([]int, m+n)
	i := 0
	k := 0
	g := 0
	for i < m && k < n {
		if nums1[i] <= nums2[k] {
			resultarray[g] = nums1[i]
			i++
		} else {
			resultarray[g] = nums2[k]
			k++
		}
		g++
	}
	//remaining
	for ; i < m; i++ {
		resultarray[g] = nums1[i]
		g++
	}
	for ; k < n; k++ {
		resultarray[g] = nums2[k]
		g++
	}

	for l := 0; l < m+n; l++ {
		nums1[l] = resultarray[l]
	}

}

func moveZerosToend(arr []int) {
	lastZeroFoundAt := 0
	for i := range arr {
		if arr[i] != 0 {
			arr[lastZeroFoundAt], arr[i] = arr[i], arr[lastZeroFoundAt]
			lastZeroFoundAt++
		}
	}
}
func TestMoveZ(t *testing.T) {
	a := []int{0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6, 0, 7, 0, 8, 0, 9, 0, 0}
	fmt.Println("Original array:", a)

	moveZerosToend(a)
	fmt.Println("Array after moving zeros to the end:", a)
}
