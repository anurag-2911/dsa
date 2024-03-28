package basics

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestXX(t *testing.T) {
	fmt.Print()
}

/*
Find the Missing Number: Given an array of n integers ranging from 1 to n+1,
with one number missing, find the missing number. The array has no duplicates.

Maximum Subarray Sum (Kadane's Algorithm): Given an array of integers,
find the contiguous subarray (containing at least one number)
which has the largest sum and return its sum.

Rotate Array: Given an array, rotate the array to the right by k steps,
where k is non-negative. Try to come up with as many solutions as you can,
and there are at least three different ways to solve this problem.

Two Sum: Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Merge Sorted Array: You are given two integer arrays nums1 and nums2,
sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
Merge nums1 and nums2 into a single array sorted in non-decreasing order.
The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
To accommodate this, nums1 has a length of m + n,
where the first m elements denote the elements that should be merged,
and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
*/

type ArrayOps struct{}

func TestMissingNum(t *testing.T) {
	tests := []struct {
		Name     string
		arr      []int
		expected int
	}{
		{"test1", []int{1, 2, 4, 5, 6, 7}, 3},
	}
	arr := &ArrayOps{}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := arr.findMissingNum(test.arr)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v and found %v ", test.expected, result)
			}
		})
	}
}
func (arr *ArrayOps) findMissingNum(array []int) int {
	n := len(array) + 1
	sumOfNNaturalNum := (n * (n + 1)) / 2
	sum := 0

	for _, val := range array {
		sum += val
	}

	return sumOfNNaturalNum - sum
}

func TestLargestSum(t *testing.T) {
	tests := []struct {
		Name     string
		Array    []int
		Expected int
	}{
		{"first", []int{-2, -3, 4, -1, -2, 1, 5, -3}, 7},
	}
	var arrops ArrayOps
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := arrops.largestSum(test.Array)
			if result != test.Expected {
				t.Errorf("expected %v and found %v", test.Expected, result)
			}
		})
	}
}
func (arops *ArrayOps) largestSum(arr []int) int {
	maxSoFar := math.MinInt32
	maxEndingHere := 0

	for _, num := range arr {
		maxEndingHere += num
		if maxEndingHere >maxSoFar {
			maxSoFar = maxEndingHere
		}
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}
	return maxSoFar
}
