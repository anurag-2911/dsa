package gobasics

import (
	"fmt"
	"testing"
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
	fmt.Println(gridTravellermemo(1, 1,make(map[string]int))) //1
	fmt.Println(gridTravellermemo(2, 3,make(map[string]int))) //3
	fmt.Println(gridTravellermemo(3, 2,make(map[string]int))) //3
	fmt.Println(gridTravellermemo(3, 3,make(map[string]int))) //6

	fmt.Println(gridTravellermemo(18, 18,make(map[string]int))) //2333606220


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
