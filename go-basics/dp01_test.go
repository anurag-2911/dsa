package gobasics

import (
	"fmt"
	"hash/fnv"
	"sort"
	"testing"
	"time"
)

func TestY(t *testing.T) {
	fmt.Println()
}

func TestFibonacci(t *testing.T) {
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fibmemo(6, make(map[int]int)))
	fmt.Println(fibmemo(7, make(map[int]int)))
	fmt.Println(fibmemo(50, make(map[int]int)))

}
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func fibmemo(n int, memo map[int]int) int {
	if val, found := memo[n]; found {
		return val
	}
	if n <= 1 {
		memo[n] = n
		return memo[n]
	}
	memo[n] = fibmemo(n-1, memo) + fibmemo(n-2, memo)
	return memo[n]

}
func gridTraveller(x, y int) int {
	if x == 1 && y == 1 {
		return 1
	}
	if x == 0 || y == 0 {
		return 0
	}
	return gridTraveller(x-1, y) + gridTraveller(x, y-1)
}
func TestGridTraveller(t *testing.T) {
	fmt.Println(gridTravellermemo(1, 1, make(map[string]int))) //1
	fmt.Println(gridTravellermemo(2, 3, make(map[string]int))) //3
	fmt.Println(gridTravellermemo(3, 2, make(map[string]int))) //3
	fmt.Println(gridTravellermemo(3, 3, make(map[string]int))) //6

	fmt.Println(gridTravellermemo(18, 18, make(map[string]int))) //2333606220

	fmt.Println(gridTraveller(1, 1)) //1
	fmt.Println(gridTraveller(2, 3)) //3
	fmt.Println(gridTraveller(3, 2)) //3
	fmt.Println(gridTraveller(3, 3)) //6

	fmt.Println(gridTraveller(18, 18)) //2333606220

}
func gridTravellermemo(n, m int, memo map[string]int) int {
	key := fmt.Sprintf("%d:%d", n, m)
	if n == 1 && m == 1 {
		memo[key] = 1
		return memo[key]
	}
	if n == 0 || m == 0 {
		memo[key] = 0
		return memo[key]
	}
	if val, found := memo[key]; found {
		return val
	}
	memo[key] = gridTravellermemo(n-1, m, memo) + gridTravellermemo(n, m-1, memo)
	return memo[key]

}

func TestCanSum(t *testing.T) {

	fmt.Println(canSummemo(7, []int{2, 3}, make(map[int]bool)))
	fmt.Println(canSummemo(7, []int{5, 3, 4, 7}, make(map[int]bool)))
	fmt.Println(canSummemo(7, []int{2, 4}, make(map[int]bool)))
	fmt.Println(canSummemo(300, []int{7, 14}, make(map[int]bool)))

	fmt.Println(canSum(7, []int{2, 3}))
	fmt.Println(canSum(7, []int{5, 3, 4, 7}))
	fmt.Println(canSum(7, []int{2, 4}))
	fmt.Println(canSum(300, []int{7, 14}))

}

// can sum to a target problem

func canSum(target int, arr []int) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}

	for i := 0; i < len(arr); i++ {
		remainder := target - arr[i]
		if canSum(remainder, arr) == true {
			return true
		}
	}
	return false
}
func canSummemo(target int, arr []int, memo map[int]bool) bool {

	if val, found := memo[target]; found {
		return val
	}

	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}

	for i := 0; i < len(arr); i++ {
		remainder := target - arr[i]
		if canSummemo(remainder, arr, memo) {
			memo[target] = true
			return true
		}
	}

	memo[target] = false
	return memo[target]
}

func TestHowSum(t *testing.T) {
	fmt.Println(howSum(7, []int{2, 3}))
	fmt.Println(howSum(7, []int{5, 3, 4, 7}))
	fmt.Println(howSum(7, []int{2, 4}))
	fmt.Println(howSum(300, []int{7, 14}))
}
func howSum(target int, arr []int) []int {
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	for _, val := range arr {
		remainder := target - val
		remainderResult := howSum(remainder, arr)
		if remainderResult != nil {
			return append(remainderResult, val)
		}
	}

	return nil
}

func howSummemo(target int, arr []int, memo map[int][]int) []int {
	if val, found := memo[target]; found {
		if val != nil {
			return append([]int(nil), val...)
		}
		return nil
	}
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	for _, val := range arr {
		remainder := target - val
		remainderResult := howSummemo(remainder, arr, memo)
		if remainderResult != nil {
			result := append(remainderResult, val)
			memo[target] = result
			return result
		}
	}
	return nil
}

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(nums)
	fmt.Println("Maximum Subarray Sum:", result) // Output: 6
}
func maxSubArray(nums []int) int {
	maxSoFar := nums[0]
	maxEndingHere := nums[0]

	for i := 1; i < len(nums); i++ {
		// Calculate max subarray sum ending at index i
		current := nums[i]
		maxEndingHere = max(current, maxEndingHere+nums[i])

		// Update max so far if needed
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestTypeAssertion(t *testing.T) {
	var g Gen = "Hallo"

	if val, ok := g.(string); ok {
		fmt.Println(val)
	}
	if val, ok := g.(int); ok {
		fmt.Println(val)
	}
	var x interface{} = 1234

	if val, ok := x.(int); ok {
		fmt.Println(val)
	}
	if val, ok := x.(string); ok {
		fmt.Println(val)
	}

}

type Gen interface{}

func TestVariadicFunctions(t *testing.T) {
	fmt.Println(sumup(1, 2))
	fmt.Println(sumup(10, 2, 1, 2))
}
func sumup(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum = sum + val
	}
	return sum
}

func TestHashTable(t *testing.T) {
	ht := NewHashTable(10)
	ht.Put("name", "John Doe")
	ht.Put("age", 30)

	if name, ok := ht.Get("name"); ok {
		fmt.Println("Name:", name)
	}

	if age, ok := ht.Get("age"); ok {
		fmt.Println("Age:", age)
	}

	ht.Remove("name")
	if name, ok := ht.Get("name"); !ok {
		fmt.Println("Name not found", name)
	}
}

type KeyValue struct {
	Key   string
	Value interface{}
}

type HashTable struct {
	buckets [][]KeyValue
	size    int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([][]KeyValue, size),
		size:    size,
	}
}
func (h *HashTable) hash(key string) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	return int(hasher.Sum32()) % h.size
}

func (h *HashTable) Put(key string, value interface{}) {
	index := h.hash(key)
	bucket := h.buckets[index]

	for i, keyvalue := range bucket {
		if keyvalue.Key == key {
			h.buckets[index][i].Value = value
			return
		}
	}
	h.buckets[index] = append(bucket, KeyValue{Key: key, Value: value})
}
func (h *HashTable) Get(key string) (interface{}, bool) {
	index := h.hash(key)
	bucket := h.buckets[index]

	for _, kv := range bucket {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return nil, false
}
func (h *HashTable) Remove(key string) {
	index := h.hash(key)
	bucket := h.buckets[index]
	for i, kv := range bucket {
		if kv.Key == key {
			h.buckets[index] = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}

func TestRateLimiter(t *testing.T) {
	action := func() {
		fmt.Println("action run at ", time.Now())
	}
	limiter := NewRateLimiter(5 * time.Second)
	limiter.Start(action)

	time.Sleep(30 * time.Second)
	limiter.Stop()

	fmt.Println("all done")
}

type RateLimiter struct {
	interval time.Duration
	ticker   *time.Ticker
	quit     chan struct{}
}

func NewRateLimiter(interval time.Duration) *RateLimiter {
	return &RateLimiter{
		interval: interval,
		ticker:   time.NewTicker(interval),
		quit:     make(chan struct{}),
	}
}
func (r *RateLimiter) Start(action func()) {
	go func() {
		for {
			select {
			case <-r.ticker.C:
				{
					action()
				}
			case <-r.quit:
				{
					r.ticker.Stop()
				}
			}
		}
	}()
}
func (r *RateLimiter) Stop() {
	close(r.quit)
}

func TestSort(t *testing.T){
	arr:=[]int{23,4,1,67,4}
	sort.Ints(arr)
	fmt.Println(arr)
	str:=[]string{"s","a","y"}
	sort.Strings(str)
	fmt.Println(str)
	people := []Person{
        {"Bob", 31},
        {"John", 42},
        {"Michael", 17},
        {"Jenny", 26},
    }
	sort.Sort(ByAge(people))
	fmt.Println(people)

}

type Person struct{
	Name string
	Age int
}

type ByAge []Person

func(a ByAge)Len()int{
	return len(a)
}
func(a ByAge)Less(i,j int)bool{
	return a[i].Age<a[j].Age
}
func(a ByAge)Swap(i,j int){
	a[i],a[j]=a[j],a[i]

}
