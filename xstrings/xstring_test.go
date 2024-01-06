package xstrings

import (
	"fmt"
	"sort"
	"testing"
)

func TestXxx(t *testing.T) {

}
func TestUniqueChars(t *testing.T) {
	fmt.Println(isUnique("Hallo"))
	fmt.Println(isUnique("Hola"))
}
func isUnique(s string) bool {
	if len(s) > 256 {
		return false
	}

	charSet := make(map[rune]bool)

	for _, char := range s {
		if _, exists := charSet[char]; exists {
			return false
		}
		charSet[char] = true

	}
	return true
}
func TestLongestSubs(t *testing.T) {
	fmt.Println(longsubs("abcabcbbxyzfgh"))
	fmt.Println(longsubs("bbbbb"))
}
func longsubs(str string) int {
	charset := make(map[rune]int)
	max := 0
	start := 0
	currentmax := 0
	for index, char := range str {
		if _, found := charset[char]; found {
			start = 0
			if currentmax > max {
				max = currentmax
			}
			currentmax = 0
		} else {
			charset[char] = index
			if start <= index {
				currentmax++
			}
		}
	}
	if currentmax > max {
		return currentmax
	}

	return max
}

func TestAnagramCheck(t *testing.T) {
	fmt.Println(isanagram("listen", "silent"))
	fmt.Println(isanagram("hello", "world"))
}
func isanagram(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	if sortString(str1) == sortString(str2) {
		return true
	}
	return false
}

func sortString(str string) string {
	r := []rune(str)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func TestPalin(t *testing.T) {
	fmt.Println(isPalin("racecar"))
	fmt.Println(isPalin("openAI"))
}
func isPalin(str1 string) bool {
	r:=[]rune(str1)
	j:=len(r)-1

	for i:=0;i<len(r)/2;i++{
		if r[i]!=r[j]{
			return false
		}
		j--
	}


	return true
}
