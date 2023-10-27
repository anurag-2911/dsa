package dynamicprogramming

import (
	"fmt"
	"testing"
)

var (
	countdb      int
	countfibmemo int
)

func TestFibo(t *testing.T) {
	fmt.Println(fibmemoized(60, make(map[int]int)))
	fmt.Println("count of fibmemo ", countfibmemo)
	dib(4)

	fmt.Println("count of db ", countdb)
	Fibo()
}
func dib(n int) {
	countdb++
	fmt.Print(n)
	fmt.Print(" | ")
	if n <= 1 {
		return
	}
	dib(n - 1)
	dib(n - 1)

}
func fibmemoized(n int, temp map[int]int) int {
	countfibmemo++
	if val, exists := temp[n]; exists {
		return val
	}
	if n <= 1 {
		return 1
	}
	temp[n] = fibmemoized(n-1, temp) + fibmemoized(n-2, temp)
	return temp[n]
}
