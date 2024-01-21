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
	fmt.Println(gridTravelermemo(2, 3,make(map[string]int)))
	fmt.Println(gridTravelermemo(3, 2,make(map[string]int)))
	fmt.Println(gridTravelermemo(3, 3,make(map[string]int)))
	fmt.Println(gridTravelermemo(30, 30,make(map[string]int)))

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

