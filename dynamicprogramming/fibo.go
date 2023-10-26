package dynamicprogramming

import (
	"fmt"
	"time"
)

var numofCalls int = 0
var numofxar int = 0

func Fibo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic ", r)
		}
	}()

	testFunction()
}

func testFunction() {
	xar(6)
	fmt.Println(numofxar)
	fmt.Println("tests")
	curentTime := time.Now()
	times:=10
	fmt.Println(fib(times))
	time.Since(curentTime).Milliseconds()
	fmt.Println("number of times called fibo", numofCalls)
	numofCalls=0
	fmt.Println(fibonacci(times,make(map[int]int)))
	fmt.Println("number of times called fibonacci",numofCalls)


	foo(4)
	fmt.Println("number of times foo called ", fooCount)
	bar(4)

}

func fib(n int) int {
	numofCalls++
	if n <= 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}

}

var fooCount int

func foo(n int) {
	fooCount++
	if n <= 1 {
		return
	}
	foo(n - 1)
}
func bar(n int) {
	if n <= 1 {
		return
	}
	bar(n - 2)
}
func xar(n int) {
	numofxar++
	if n <= 1 {
		return
	}
	xar(n - 1)
	xar(n - 1)
}
func fibonacci(n int, memo map[int]int) int {
	numofCalls++
	//check if the value is already calculated
	if val, found := memo[n]; found {
		return val
	}
	//base case
	if n <= 1 {
		memo[n] = n
		return 1
	}
	memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
	return memo[n]

}
