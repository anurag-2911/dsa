package util

import
(

)

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
