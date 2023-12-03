package arrays

import (
	"fmt"
	"math"
	"testing"
)

func TestXxx(t *testing.T) {

}

func findMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}
func TestFindMax(t *testing.T) {
	input := []int{34, 5, 67, 8, 9}
	fmt.Println(findMax(input))
}
func findSecondMax(arr []int) int {
	max, secmax := math.MinInt64, math.MinInt64

	for _, num := range arr {
		if num > max {
			secmax = max
			max = num
		} else if num > secmax && num != max {
			secmax = num
		}
	}
	return secmax
}
func TestSecondMax(t *testing.T) {
	input := []int{34, 5, 67, 8, 9}
	fmt.Println(findSecondMax(input))
	fmt.Println(findSecondMax([]int{12, 35, 1, 10, 34, 1}))
}
func checkifarraySorted(arr []int) bool {
	initial := math.MinInt64
	for _, num := range arr {
		if num >= initial {
			initial = num
		} else {
			return false
		}
	}
	return true
}
func TestCheckArraySorted(t *testing.T) {
	fmt.Println(checkifarraySorted([]int{1, 3, 5, 7, 9, 11}))
	fmt.Println(checkifarraySorted([]int{1, 1, 1, 1, 1, 1}))
	fmt.Println(checkifarraySorted([]int{1, 3, 15, 7, 9, 11}))

}

func TestArrayRotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	rotate := 2
	// result:=[]int{3,4, 5, 1, 2}
	res := rotateArray(arr, rotate)
	fmt.Println(res)

}
func rotateArray(arr []int, rotateby int) []int {
	nearr := make([]int, 0, len(arr))

	for i := 0; i < len(arr); i++ {
		index := (i + rotateby) % len(arr)
		nearr = append(nearr, arr[index])
	}
	return nearr
}
func TestFindMissingNumber(t *testing.T) {
	arr := []int{1, 2, 4, 5, 6}
	fmt.Println(findMissingNum(arr))
	fmt.Println(findMissingNum([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(findMissingNum([]int{1, 2, 3, 4, 5, 7}))
	fmt.Println(findMissingNum([]int{}))
}
func findMissingNum(arr []int) int {
	if len(arr) <= 0 {
		return -1
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]+1 != arr[i+1] {
			return arr[i] + 1
		}
	}
	return -1
}

func mergesortedarrays(arr1 []int, arr2 []int) []int {
	i, j := 0, 0
	result := make([]int, 0, len(arr1)+len(arr2))

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}

	}

	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)

	return result
}
func TestMergeSortedArray(t *testing.T) {
	fmt.Println(mergesortedarrays([]int{1, 3, 5}, []int{2, 4, 6}))
}
