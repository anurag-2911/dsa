package recursion

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	// result := A()
	// fmt.Println(result)
	fmt.Println(revstring("halo"))
}

func revstring(s string) string {
	if len(s) == 0 {
		return s
	}
	substr := s[1:]
	first := s[0]

	return revstring(substr) + string(first)
}

func TestPalinDrome(t *testing.T) {
	fmt.Println(isPalindrome("racecar"))
}

func isPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	return checkPalindrome(s, 0, len(s)-1)

}
func checkPalindrome(s string, start, end int) bool {
	if start >= end {
		return true
	}
	if s[start] != s[end] {
		return false
	}
	return checkPalindrome(s, start+1, end-1)
}
func checkPalin(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	if s[0] == s[len(s)-1] {
		return checkPalin(s[1 : len(s)-1])
	}
	return false
}
func TestCheckPalin(t *testing.T) {
	fmt.Println(checkPalin("racecar"))
	fmt.Println(checkPalin("halkah"))
}
func decimaltoBinary(n int) string {
	if n == 0 {
		return ""
	}
	return decimaltoBinary(n/2) + strconv.Itoa(n%2)
}
func TestDecToBinary(t *testing.T) {
	fmt.Println(decimaltoBinary(7))
}
