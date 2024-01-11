package basicprogs

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	fmt.Print()
}

// basics programming questions

// fibonacci series

func TestFibo(t *testing.T) {
	fmt.Println(fibomem(5, make(map[int]int)))
	fmt.Println(fibomem(6, make(map[int]int)))
	fmt.Println(fibomem(7, make(map[int]int)))
	fmt.Println(fibomem(50, make(map[int]int)))

	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(50))
}

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func fibomem(n int, memo map[int]int) int {
	if val, exists := memo[n]; exists {
		return val
	}
	if n <= 1 {
		return n
	}
	memo[n] = fibomem(n-1, memo) + fibomem(n-2, memo)
	return memo[n]

}

// grid traveller problem

func TestGridTraveller(t *testing.T) {
	fmt.Println(gridTravellermemo(1, 1,make(map[string]int)))
	fmt.Println(gridTravellermemo(3, 2,make(map[string]int)))
	fmt.Println(gridTravellermemo(2, 3,make(map[string]int)))
	fmt.Println(gridTravellermemo(3, 3,make(map[string]int)))
	fmt.Println(gridTravellermemo(18, 18,make(map[string]int)))
	fmt.Println(gridTraveller(1, 1))
	fmt.Println(gridTraveller(3, 2))
	fmt.Println(gridTraveller(2, 3))
	fmt.Println(gridTraveller(3, 3))
	fmt.Println(gridTraveller(18, 18))
}
func gridTraveller(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}
	if x == 1 && y == 1 {
		return 1
	}
	return gridTraveller(x-1, y) + gridTraveller(x, y-1)
}

func gridTravellermemo(x, y int, memo map[string]int) int {
	key := fmt.Sprintf("%d:%d", x, y)
	if val, found := memo[key]; found {
		return val
	}
	if x==0 || y==0{
		return 0
	}
	if x==1 && y==1{
		memo[key]=1
		return memo[key]
	}
	memo[key] = gridTravellermemo(x-1, y, memo) + gridTravellermemo(x, y-1, memo)
	return memo[key]
}
