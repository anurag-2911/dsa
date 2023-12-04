package dynamicprogramming

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println()
}

func TestXFibo(t *testing.T) {
	//1,1,2,3,5,8,13,21,34....
	fmt.Println(xfibmemo(6,make(map[int]int)))
	fmt.Println(zoo(5))
	fmt.Println(fibx(6))
}

// 2^n time complexity and n space (stack frames at any point of time) complexity
func fibx(n int) int {
	if n <= 2 {
		return 1
	}
	n1 := fibx(n - 1)
	n2 := fibx(n - 2)
	return n1 + n2
}

func zoo(n int) int {
	if n <= 1 {
		return 1
	}
	return zoo(n - 1)
}
func xfibmemo(n int, result map[int]int) int {
	if v, exists := result[n]; exists {
		return v
	}
	if n <= 2 {
		return 1
	}
	res := xfibmemo(n-1, result) + xfibmemo(n-2, result)
	result[n] = res
	return res
}
