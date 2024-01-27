package basics

import (
	"fmt"
	"reflect"
	"strings"
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

//end of double linked list

/*
strings problems
string reversal
*/
func strrev(str string) string {
	r := make([]rune, 0)
	for i := len(str) - 1; i >= 0; i-- {
		r = append(r, rune(str[i]))
	}
	return string(r)
}
func strrev01(str string) string {
	r := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func TestStringReverse(t *testing.T) {
	fmt.Println(strrev("hallo"))
	fmt.Println(strrev01("hallo"))
}
func palin(str string) bool {
	r := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}
func TestPalin(t *testing.T) {
	fmt.Println(palin("racecar"))
	fmt.Println(palin("sidecar"))
	fmt.Println(palin("raccar"))
	fmt.Println(palin("racxxar"))
}

func isAnagram(str1, str2 string) bool {
	lstr1 := strings.ToLower(str1)
	lstr2 := strings.ToLower(str2)
	if len(str1) != len(str2) {
		return false
	}
	frequency := make(map[rune]int)
	for _, v := range lstr1 {
		frequency[v]++
	}
	for _, v := range lstr2 {
		frequency[v]--
		if frequency[v] < 0 {
			return false
		}
	}
	return true
}
func TestAnagram(t *testing.T) {
	fmt.Println(isAnagram("Listen", "Silent"))     // Should return true
	fmt.Println(isAnagram("Triangle", "Integral")) // Should return true
	fmt.Println(isAnagram("Hello", "World"))       //false
}
func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	//abcabcbb
}
func longestsubstringwithoutrepeatingcharacters(str string) int {
	return 0

}

//end if strings problems

/*
Arrays problems

*/

// array rotation
type Arrays struct {
}

var arrays Arrays = Arrays{}

func TestArrays(t *testing.T) {
	a := Arrays{}
	got := a.rotate([]int{1, 2, 3, 4, 5}, 2)
	want := []int{3, 4, 5, 1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v,want %v", got, want)
	}
}
func (a *Arrays) rotate(arr []int, rotateby int) []int {
	rotatedarr := make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		index := (i + rotateby) % len(arr)
		rotatedarr = append(rotatedarr, arr[index])
	}
	return rotatedarr
}

/*
Find the Missing Number:

Description: You are given an array of n integers from 1 to n+1,
			 such that all numbers from 1 to n+1 are present except one missing number.
*/

func TestMissingNumber(t *testing.T) {
	a := Arrays{}
	got := a.findMissingnum([]int{1, 2, 3, 4, 6, 7, 8})
	want := 5

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d,want %d", got, want)
	}

}
func (a *Arrays) findMissingnum(arr []int) int {
	n := len(arr)
	total := ((n + 1) * (n + 2)) / 2
	sum := 0
	for _, val := range arr {
		sum = sum + val
	}
	return total - sum
}

/*
Merge Sorted Arrays:

Description: Write a Go function that merges two sorted arrays into one larger sorted array.
For example, merging [1, 3, 5] and [2, 4, 6] should yield [1, 2, 3, 4, 5, 6].
Key Concepts: Two-pointer technique, understanding of sorted data structures.
*/
func (a *Arrays) merger(arr1, arr2 []int) []int {
	resultarray := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			resultarray = append(resultarray, arr1[i])
			i++
		} else {
			resultarray = append(resultarray, arr2[j])
			j++
		}
	}
	resultarray = append(resultarray, arr1[i:]...)
	resultarray = append(resultarray, arr2[j:]...)
	
	return resultarray
}
func TestMerge(t *testing.T) {
	a := Arrays{}
	got := a.merger([]int{1, 3, 5}, []int{2, 4, 6})
	want := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

//end of arrays problems
