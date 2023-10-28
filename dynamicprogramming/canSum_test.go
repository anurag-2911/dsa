package dynamicprogramming

import (
	printme "dsa/helper"
	"fmt"
	"testing"
)

func TestCanSum(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from a panic", r)
		}
	}()
	runTests()
}

func runTests() {
	arr := []int{5, 3, 4, 7}
	result := canSum(arr, 7,make(map[int]bool))
	fmt.Println(result)
}

func canSum(arr []int, sum int, eval map[int]bool) bool {
	printme.PrintFormatted(sum)
	if val,exists:=eval[sum];exists!=false{
		return val
	}
	if sum == 0 {
		return true
	}
	if sum < 0 {
		return false
	}
	for i := 0; i < len(arr); i++ {
		reduce := sum - arr[i]
		if canSum(arr, reduce, eval) {
			eval[sum] = true
			return eval[sum]
		}
	}

	eval[sum] = false
	return eval[sum]
}
