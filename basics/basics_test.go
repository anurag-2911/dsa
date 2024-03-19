package basics

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestX(t *testing.T) {
	fmt.Print()
}

/*
Write a Go function that takes a string as input and reverses the order of words
(while preserving the order of characters within each word).
Example: "Hello world today" -> "today world Hello"
*/

func revwords(words string) string {
	result := strings.Fields(words)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return strings.Join(result, " ")
}
func TestRevWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world today", "today world hello"},
		{"this is a test", "test a is this"},
		{"", ""},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			output := revwords(tc.input)
			if output != tc.expected {
				t.Errorf("expected %v found %v", tc.expected, output)
			}
		})
	}

}

/*
Implement a basic stack data structure in Go. Provide functions for the following operations:
Push(value)
Pop()
IsEmpty()
*/

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, fmt.Errorf("stack is empty")
	}
	lastindex := len(s.items) - 1
	lastItem := s.items[lastindex]
	s.items = s.items[0:lastindex]
	return lastItem, nil
}
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
func TestStackOps(t *testing.T) {
	s := &Stack{items: make([]int, 0)}
	s.Push(1000)
	s.Push(2000)
	s.Push(3000)
	fmt.Println(s.Pop())

}

/*
You have a list of URLs. Write a Go program that uses goroutines to fetch the content of each URL concurrently.
Limit the number of simultaneous goroutines to a reasonable value (e.g., 2).
*/

type URLContents struct {
	urls []string
}

func TestULRS(t *testing.T) {
	urlfetch := &URLContents{urls: make([]string, 0)}
	urlfetch.urls = append(urlfetch.urls, "https://www.google.com/")
	urlfetch.urls = append(urlfetch.urls, "https://www.yahoo.com/")
	urlfetch.urls = append(urlfetch.urls, "https://www.facebook.com/")
	urlfetch.urls = append(urlfetch.urls, "https://about.meta.com/")
	urlfetch.urls = append(urlfetch.urls, "https://www.instagram.com/")
	urlfetch.urls = append(urlfetch.urls, "https://www.google.com/maps")

	maxConcurrent := 2

	var wg sync.WaitGroup

	semaphore := make(chan struct{}, maxConcurrent)

	for _, url := range urlfetch.urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			semaphore <- struct{}{}

			//http request
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("error fetching %s url %v\n", url, err)
				return
			}
			defer resp.Body.Close()
			content, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("error in reading the content for url %s\n %v ", url, err)
				return
			}
			fmt.Println("content for url", url, string(content))
			<-semaphore //release the semaphore
		}(url)
	}
	wg.Wait()

}

/*
Write a Go function that opens a file, reads its contents, and processes the data.
Ensure robust error handling throughout the process
(e.g., gracefully handle cases where the file doesn't exist,
or there are errors during the reading or processing).
*/

func TestFileOps(t *testing.T) {
	fmt.Println(fileops("./basics_test.go"))
}
func fileops(filePath string) (string, error) {
	fs, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer fs.Close()
	data, err := io.ReadAll(fs)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
go routines and channels
*/
func TestGoRoutines(t *testing.T) {
	var wg sync.WaitGroup

	genGoRoutines(&wg)
	wg.Wait()
	fmt.Println("all done")

}
func genGoRoutines(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("I am routine %d for %d\n counter", rand.Intn(i+100), i)
		}(i)
	}
}

/*
Problem 2:  Shared Data (with Race Conditions)

Create a program with multiple goroutines incrementing a global counter variable.
Observe the race condition that occurs.
*/

var counter int = 100

func incCount() {
	for i := 0; i < 5; i++ {
		counter++
	}
}
func decCount() {
	for i := 0; i < 5; i++ {
		counter--
	}
}
func TestRaceCond(t *testing.T) {
	go incCount()
	go decCount()
	// time.Sleep(time.Second)
	if counter != 100 {
		t.Errorf("expected %d, found %d", 100, counter)
	}
	fmt.Println("all done")
}

// arrays common problems

type XArray struct{}

func TestArr(t *testing.T) {
	xarr := &XArray{}
	found := xarr.missingnum([]int{1, 2, 3, 5, 6, 7, 8})
	expected := 4
	if found != expected {
		t.Errorf("found %d,expected %d", found, expected)
	}

}

/*
Given an array of unique integers from 1 to n with one number missing,
write a function to find the missing number.
Consider the array is not sorted and has no duplicates.
*/
func (xa *XArray) missingnum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	n := len(arr) + 1
	sumofn := (n * (n + 1)) / 2
	diff := sumofn - sum
	return diff
}

/*
Implement a function that merges two sorted arrays into a single sorted array.
Assume both input arrays are sorted in non-decreasing order
*/

func mergethem(arr1 []int, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
		if i < len(arr1) {
			result = append(result, arr1[i:]...)
		}
		if j < len(arr2) {
			result = append(result, arr2[j:]...)
		}

	}
	return result
}
func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected []int
	}{
		{"both Empty", []int{}, []int{}, []int{}},
		{"first emptry", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := mergethem(tc.arr1, tc.arr2)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v and found %v for %s", tc.expected, result, tc.name)
			}
		})
	}
}

/*
	Write a function that determines if two strings are anagrams of each other.
	An anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
	typically using all the original letters exactly once
*/

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	wordlist := make(map[rune]int)

	// count occurences of each rune in s1
	for _, v := range s1 {
		wordlist[v]++
	}

	for _, v := range s2 {
		if _, exists := wordlist[v]; exists {
			wordlist[v]--
			if wordlist[v] == 0 {
				delete(wordlist, v)
			}
		} else {
			return false
		}

	}

	return len(wordlist) == 0
}

func TestAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{"valid", "secure", "rescue", true},
		{"invalid", "secure", "ressue", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isAnagram(test.s1, test.s2)
			if result != test.expected {
				t.Errorf("expected %v and got %v ", test.expected, result)
			}
		})
	}
}

/*
Implement a function that checks if a string is a palindrome.
A palindrome is a word, phrase, number, or other sequences of characters which reads the same backward
as forward, ignoring spaces, punctuation, and case sensitivity.
*/

func isPalindrome(word string) bool {

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}

	return true
}
func TestPalidrome(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		expected bool
	}{
		{"valid", "racecar", true},
		{"invalid", "sidecar", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			found := isPalindrome(test.word)
			if found != test.expected {
				t.Errorf("expected %v and found %v for test %v", test.expected, found, test.name)
			}
		})
	}
}

/*
Write a function that finds the length of the longest substring without repeating characters.
For example, the longest substring without repeating letters for "abcabcbb" is "abc", which the length is 3.
*/

func longestsubs(word string) int {
	wordlist := make(map[rune]int)

	for _, w := range word {
		if _, exists := wordlist[w]; !exists {
			wordlist[w]++
		}
	}
	return len(wordlist)
}

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		expected int
	}{
		{"first", "abcabcbb", 3}, {"second", "abcdefda", 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := longestsubs(test.word)
			if result != test.expected {
				t.Errorf("expected %v and found %v ,for test %v", test.expected, result, test.name)
			}
		})
	}
}

/*
Implement a method to perform basic string compression using the counts of repeated characters.
For example, the string "aabcccccaaa" would become "a2b1c5a3".
If the "compressed" string would not become smaller than the original string,
your function should return the original string
*/
func compresswords(word string) string {
	var result strings.Builder
	count := 1
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			count++
		} else {
			result.WriteByte(word[i-1])
			result.WriteString(strconv.Itoa(count))
			count = 1
		}
	}
	// add the last character and its count
	result.WriteByte(word[len(word)-1])
	result.WriteString(strconv.Itoa(count))
	compressed := result.String()
	if len(compressed) >= len(word) {
		return word
	}
	return compressed
}

func TestCompressedSrting(t *testing.T) {
	tests := []struct {
		Name         string
		Word         string
		ExpectedWord string
	}{
		{"SimpleTest", "aabcccccaaa", "a2b1c5a3"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := compresswords(test.Word)
			if result != test.ExpectedWord {
				t.Errorf("expected %v,found %v for %v", test.ExpectedWord, result, test.Name)
			}
		})
	}
}

type StringOps struct{}

func (sops *StringOps) palinchecker(word string) bool {
	length := len(word) - 1

	for i, j := 0, length; i < length; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}
func TestStringsOps(t *testing.T) {
	tests := []struct {
		Name     string
		Word     string
		Expected bool
	}{
		{"racecar", "racecar", true},
		{"race", "race", false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			sops := &StringOps{}
			result := sops.palinchecker(test.Word)
			if result != test.Expected {
				t.Errorf("expected %v for %v got %v ", test.Expected, test.Name, result)
			}
		})
	}
}
func TestAnag(t *testing.T) {
	tests := []struct {
		Name     string
		Word1    string
		Word2    string
		Expected bool
	}{
		{"positive", "worth", "throw", true},
	}
	for _, test := range tests {
		sops := &StringOps{}
		result := sops.isAnagram(test.Word1, test.Word2)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("expected %v,found %v", test.Word1, result)
		}
	}
}
func (sops *StringOps) isAnagram(word1, word2 string) bool {
	wordlist := make(map[rune]int)

	for _, v := range word1 {
		wordlist[v]++
	}

	for _, v := range word2 {
		if _, exists := wordlist[v]; exists {
			wordlist[v]--
			if wordlist[v] == 0 {
				delete(wordlist, v)
			}
		}
	}

	return len(wordlist) == 0
}

/*
Implement a method in Go to perform basic string compression using the counts of repeated characters.
For example, the string "aabcccccaaa" would become "a2b1c5a3".
If the "compressed" string would not become smaller than the original string,
your program should return the original string
*/

func TestStringCompress(t *testing.T) {
	tests := []struct {
		TestName string
		Word     string
		Expected string
	}{
		{"first", "aabcccccaaa", "a2b1c5a3"},
	}
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			sops := &StringOps{}
			result := sops.compressed(test.Word)
			if !reflect.DeepEqual(result, test.Expected) {
				t.Errorf("wrong")
			}
		})
	}

}
func (sops *StringOps) compressed(word string) string {
	count := 1
	var sb strings.Builder
	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			count++
		} else {
			sb.WriteByte(word[i-1])
			sb.WriteString(strconv.Itoa(count))
			count = 1
		}
	}
	sb.WriteByte(word[len(word)-1])
	sb.WriteString(strconv.Itoa(count))
	return sb.String()
}

/*
Write a Go program that takes a string containing any combination of the characters
'(', ')', '{', '}', '[' and ']', and checks if the brackets are balanced.
For brackets to be balanced, every opening bracket must have a corresponding
closing bracket of the same type, and brackets must close in the correct order

*/

func TestBrackets(t *testing.T) {
	tests := []struct {
		Name       string
		Word       string
		IsBalanced bool
	}{
		{"first", "(xcf)", true},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			sops := &StringOps{}
			result := sops.IsBalanced(test.Word)
			if !reflect.DeepEqual(result, test.IsBalanced) {
				t.Errorf("wrong!!")
			}
		})
	}
}

func TestProperBracket(t *testing.T){
	// Test cases
    tests := []string{"()", "()[]{}", "(]", "([)]", "{[]}", "[({})]", ""}
	sop:=&StringOps{}
    for _, test := range tests {
        fmt.Printf("%s: %t\n", test, sop.IsBalanced(test))
    }
}
func (sop *StringOps) IsBalanced(word string) bool {
	//stack to hold opening bracket
	var stack []rune

	//map to hold matching opening bracket to closing bracket

	bracketPairs := map[rune]rune{')': '(', ']': '[', '{': '}'}

	for _, char := range word {
		switch char {
		case '(', '{', '[':
			{
				stack = append(stack, char)
			}
		case ')', '}', ']':
			{
				if len(stack)==0 {
					return false
				}
				topelem:=stack[len(stack)-1]
				stack=stack[0:len(stack)-1]
				if bracketPairs[char]!=topelem{
					return false
				}
			}

		}
	}

	return len(stack) == 0
}

/*
	Implement a function in Go that finds the length of the longest substring without repeating characters. 
	For example, given "abcabcbb", the answer is "abc", which the length is 3. For "bbbbb" the longest substring is "b", 
	with the length of 1
*/

func TestLongestSubs(t *testing.T){
	tests:=[]string{"abcabcbb","bbbbb"}
	for _,test:=range tests{
		sops:=&StringOps{}
		result:=sops.findLongestSubsLen(test)
		fmt.Print(result)
		fmt.Print(" ")
	}
}
func(sops *StringOps)findLongestSubsLen(word string)int{
	count:=0
	charList:=make(map[rune]int)

	for _,char:=range word{
		if _,exists:=charList[char];exists{
			break
		}else{
			charList[char]++
			count++
		}
		
	}
	return count
}
