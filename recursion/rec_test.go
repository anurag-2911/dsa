package recursion

import (
	"fmt"
	"strconv"
	"testing"
)

func TestZXxx(t *testing.T) {

}

func xrevstring(s string) string {
	if len(s) == 0 {
		return ""
	}
	substr := s[1:]
	first := string(s[0])
	return xrevstring(substr) + first

}
func TestRevString(t *testing.T) {
	fmt.Println(xrevstring("hallo"))
}

func xpalin(s string) bool {

	if len(s) == 0 || len(s) == 1 {
		return true
	}
	firstChar := s[0]
	lastchar := s[len(s)-1]
	if firstChar == lastchar {
		substr := s[1 : len(s)-1]
		return xpalin(substr)

	}

	return false
}
func TestXPlain(t *testing.T) {
	fmt.Println(xpalin("racecar"))
	fmt.Println(xpalin("racegurram"))
}

func xdectobin(n int) string {
	if n == 0 {
		return ""
	}
	q := n / 2
	r := n % 2
	return xdectobin(q) + strconv.Itoa(r)
}

func TestDecBin(t *testing.T) {
	fmt.Println(xdectobin(16))
	fmt.Println(xdectobin(9))

}

func sumofNaturalNums(n int) int {
	if n == 1 {
		return 1
	}

	sum := n + sumofNaturalNums(n-1)
	return sum
}

func TestSumOfNatural(t *testing.T) {
	fmt.Println(sumofNaturalNums(10))
}

func xbinarysrch(n []int, search int) bool {
	if(len(n)==0){
		return false
	}
	if len(n) == 1 && n[0] != search {
		return false
	}
	right := len(n) - 1
	left := 0
	mid := (left + right) / 2
	midnum := n[mid]
	if midnum == search {
		return true
	}
	if search < n[mid] {
		right = mid
		mid = (left + right) / 2
		return xbinarysrch(n[0:mid], search)
	} else {
		left = mid
		mid = (left + right) / 2
		return xbinarysrch(n[mid:right], search)
	}
	

}
func TestXBinary(t *testing.T) {
	fmt.Println(xbinarysrch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 8))
	fmt.Println(xbinarysrch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 14))
	fmt.Println(xbinarysrch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16}, 10))
	fmt.Println(xbinarysrch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 18))
	fmt.Println(xbinarysrch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 0))
}

func xfibonacci(n int)(int){
	if(n==0 || n==1){
		return n
	}
	n1:=n-1
	n2:=n-2
	return xfibonacci(n1)+ xfibonacci(n2)
}
func TestFibo(t *testing.T){
	//0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144
	fmt.Println(xfibonacci(5))
	fmt.Println(xfibonacci(6))
	fmt.Println(xfibonacci(7))
	fmt.Println(xfibonacci(8))


}

