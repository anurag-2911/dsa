package mergesort

import (
	"fmt"
)

var counter int = 0

func TestMergeSort() {
	fmt.Println("merge sort")
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array:", arr)
	sortedArr := Mergesort(arr)
	fmt.Println("Sorted array:", sortedArr)
	fmt.Println("number of times merge is called ", counter)

}

func Mergesort(arr []int) []int {
	counter++
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	leftArray := arr[:middle]
	rightArray := arr[middle:]

	leftSortedArry := Mergesort(leftArray)
	rightSortedArray := Mergesort(rightArray)

	return merge(leftSortedArry, rightSortedArray)
}

func merge(leftArray []int, rightArray []int) []int {
	result := make([]int, 0, len(leftArray)+len(rightArray))
	i, j := 0, 0
	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] <= rightArray[j] {
			result = append(result, leftArray[i])
			i++
		} else {
			result = append(result, rightArray[j])
			j++
		}

	}
	//remaining entries
	for ; i < len(leftArray); i++ {
		result = append(result, leftArray[i])
	}
	for ; j < len(rightArray); j++ {
		result = append(result, rightArray[j])
	}
	return result
}
