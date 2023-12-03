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
func findSecondMax(arr []int)int{
	max,secmax:=math.MinInt64,math.MinInt64

	for _,num:=range arr{
		if num>max{
			secmax=max
			max=num
		}else if num>secmax && num!=max{
			secmax=num
		}
	}
	return secmax
}
func TestSecondMax(t *testing.T){
	input := []int{34, 5, 67, 8, 9}
	fmt.Println(findSecondMax(input))
	fmt.Println(findSecondMax([]int{12, 35, 1, 10, 34, 1}))
}
