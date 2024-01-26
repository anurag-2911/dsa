package basics

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	fmt.Println()
}

/*
Fibonacci series: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144
*/

func TestFib(t *testing.T) {
	// memoized fibo recursive algo
	fmt.Println(fibmemo(5, make(map[int]int)))
	fmt.Println(fibmemo(6, make(map[int]int)))
	fmt.Println(fibmemo(50, make(map[int]int)))

	// simple fibo recursive algo
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(50))

}

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}
func fibmemo(n int, memo map[int]int) int {
	if val, found := memo[n]; found {
		return val
	}
	if n <= 1 {
		return n
	}
	memo[n] = fibmemo(n-1, memo) + fibmemo(n-2, memo)
	return memo[n]
}

/*
Grid traveler problem
*/

func TestGridTraveler(t *testing.T) {
	fmt.Println(gridTravelermemo(2, 3, make(map[string]int)))
	fmt.Println(gridTravelermemo(3, 2, make(map[string]int)))
	fmt.Println(gridTravelermemo(3, 3, make(map[string]int)))
	fmt.Println(gridTravelermemo(30, 30, make(map[string]int)))

	fmt.Println(gridTraveler(2, 3))
	fmt.Println(gridTraveler(3, 2))
	fmt.Println(gridTraveler(3, 3))
	fmt.Println(gridTraveler(30, 30))
}
func gridTraveler(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}
	if x == 1 && y == 1 {
		return 1
	}
	return gridTraveler(x-1, y) + gridTraveler(x, y-1)
}
func gridTravelermemo(x, y int, memo map[string]int) int {
	key := fmt.Sprintf("%d:%d", x, y)
	if val, found := memo[key]; found {
		return val
	}
	if x == 0 || y == 0 {
		memo[key] = 0
		return memo[key]
	}
	if x == 1 && y == 1 {
		memo[key] = 1
		return memo[key]
	}

	memo[key] = gridTravelermemo(x-1, y, memo) + gridTravelermemo(x, y-1, memo)
	return memo[key]
}

/*
Stack as an array
*/
type XStack struct {
	items []int
}

func (st *XStack) Push(item int) {
	st.items = append(st.items, item)
}
func (stack *XStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, fmt.Errorf("empty stack")
	}
	length := len(stack.items)

	item, items := stack.items[length-1], stack.items[0:length-1]

	stack.items = items

	return item, nil
}
func (stack *XStack) Peek() (int, error) {
	if stack.IsEmpty() {
		return 0, fmt.Errorf("empty stack")
	}
	length := len(stack.items)
	return stack.items[length-1], nil
}
func (stack *XStack) IsEmpty() bool {
	return len(stack.items) == 0
}
func TestStackOps(t *testing.T) {
	stack := &XStack{
		items: make([]int, 0),
	}
	stack.Push(101)
	stack.Push(102)
	stack.Push(103)
	stack.Push(104)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
	fmt.Println(stack.IsEmpty())
}

/*
Queue as an array
*/
type XQueue struct {
	items []int
}

func (q *XQueue) Enqueue(item int) {
	q.items = append(q.items, item)
}
func (q *XQueue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("empty queue")
	}
	item, items := q.items[0], q.items[1:]
	q.items = items
	return item, nil
}
func (q *XQueue) Peek() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("empty queue")
	}
	return q.items[0], nil
}
func TestQ(t *testing.T) {
	queue := &XQueue{
		items: make([]int, 0),
	}
	queue.Enqueue(101)
	queue.Enqueue(102)
	queue.Enqueue(103)
	queue.Enqueue(104)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Peek())
}

/*
LinkedList
*/
type YNode struct {
	item int
	next *YNode
}
type YLinkedList struct {
	head *YNode
}

func (list *YLinkedList) Insert(item int) {
	node := &YNode{
		item: item,
	}
	if list.head == nil {
		list.head = node
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = node

}
func (list *YLinkedList) Traverse() {
	current := list.head
	if current == nil {
		return
	}
	for current != nil {
		fmt.Print(current.item)
		fmt.Print(">")
		current = current.next
	}
}
func (list *YLinkedList) Remove(item int) {
	current := list.head
	if current == nil {
		return
	}
	if list.head.item == item {
		list.head = list.head.next
		return
	}
	prev := current
	for current.next != nil {

		if current.item == item {
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}
	if current.item == item {
		prev.next = nil
		return
	}
}
func TestLL(t *testing.T) {
	linkedlist := &YLinkedList{}
	linkedlist.Insert(1000)
	linkedlist.Insert(2000)
	linkedlist.Insert(3000)
	linkedlist.Insert(4000)
	linkedlist.Insert(5000)
	linkedlist.Insert(6000)

	linkedlist.Traverse()
	fmt.Println()
	linkedlist.Remove(2000)
	linkedlist.Traverse()
	fmt.Println()
	linkedlist.Remove(1000)
	linkedlist.Traverse()
	fmt.Println()
	linkedlist.Remove(6000)
	linkedlist.Traverse()
	fmt.Println()
	linkedlist.Remove(8000)
}

/*
Double linked list
*/
type DNode struct {
	data       int
	next, prev *DNode
}
type XDoubleLinkedList struct {
	head, tail *DNode
}

func (dll *XDoubleLinkedList) Insert(data int) {
	node := &DNode{
		data: data,
	}
	if dll.head == nil {
		dll.head = node
		dll.tail = node
		return
	}
	current := dll.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
	node.prev = current
	dll.tail = node

}
func (dll *XDoubleLinkedList) Traverse() {
	current := dll.head
	if current == nil {
		return
	}
	for current != nil {
		fmt.Print(current.data)
		fmt.Print("  ")
		current = current.next
	}
	fmt.Println()
}
func (dll *XDoubleLinkedList) TraverseReverse() {
	current := dll.tail
	if current == nil {
		return
	}
	for current != nil {
		fmt.Print(current.data)
		fmt.Print("  ")
		current = current.prev
	}
	fmt.Println()
}
func (dll *XDoubleLinkedList) Delete(data int) {
	current := dll.head
	if current.data == data { // if the first node itself
		dll.head = dll.head.next
		dll.head.prev = nil

		return
	}
	if dll.tail.data == data { // if the last node
		prev := dll.tail.prev
		prev.next = nil
		dll.tail = prev
		return
	}

	for current != nil {
		if current.data == data {
			current.next.prev = current.prev
			current.prev.next = current.next
			return
		}

		current = current.next
	}

}
func TestDoubleLinkedList(t *testing.T) {
	dll := &XDoubleLinkedList{}
	dll.Insert(1000)
	dll.Insert(2000)
	dll.Insert(3000)
	dll.Insert(4000)
	dll.Insert(5000)
	traversex(dll)
	dll.Delete(1000)
	traversex(dll)
	dll.Delete(5000)
	traversex(dll)
	dll.Insert(6000)
	dll.Insert(7000)
	dll.Insert(8000)
	traversex(dll)
	dll.Delete(6000)
	traversex(dll)

}

func traversex(dll *XDoubleLinkedList) {
	dll.Traverse()
	dll.TraverseReverse()
}
