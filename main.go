package main

import (
	"fmt"
)

/*
Given two sorted lists of integers. Implement a function that merges them such that the resulting list is still sorted.
For example, considering a list A of [2, 5, 6, 9, 11]
and a list B of [1, 1, 3, 5, 8],
the resulting list should be [1, 1, 2, 3, 5, 5, 6, 8, 9, 11].
*/
func main() {
	fmt.Println()
	result:=merge([]int{2, 5, 6, 9, 11},[]int{1, 1, 3, 5, 8})
	fmt.Println(result)
}

func merge(arr1 []int, arr2 []int) []int {
	resultarray := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			resultarray = append(resultarray, arr1[i])
			i++
		} else {
			resultarray = append(resultarray, arr2[j])
			j++
		}

	}
	resultarray = append(resultarray, arr1[i:]...)
	resultarray = append(resultarray, arr2[j:]...)

	return resultarray
}
