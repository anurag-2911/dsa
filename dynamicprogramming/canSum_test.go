package dynamicprogramming

import (
	"dsa/helper"
	"fmt"
	"testing"
)

func TestCanSum(t *testing.T) {
	arr := []int{5, 3, 4, 7}
	result := canSum(arr, 7)
	fmt.Println(result)
}

func canSum(arr []int, sum int) bool {
	helper.PrintFormatted(sum)
	if sum == 0 {
		return true
	}
	if sum < 0 {
		return false
	}
	for i := 0; i < len(arr); i++ {
		reduce := sum - arr[i]
		if canSum(arr, reduce) {
			return true
		}
	}

	return false
}
