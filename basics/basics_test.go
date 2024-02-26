package basics

import (
	"fmt"
	"strings"
	"testing"
	"net/http"
	"github.com/gin-gonic/gin"
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

type Stack struct{
	items []int

}
func(s *Stack)Push(value int){
	s.items=append(s.items,value )
}
func(s *Stack)Pop()(int,error){
	if s.IsEmpty(){
		return -1,fmt.Errorf("stack is empty")
	}
	lastindex:=len(s.items)-1
	lastItem:=s.items[lastindex]
	s.items=s.items[0:lastindex]
	return lastItem,nil
}
func(s *Stack)IsEmpty()bool{
	return len(s.items)==0
}
func TestStackOps(t *testing.T){
	s:=&Stack{items:make([]int, 0)}
	s.Push(1000)
	s.Push(2000)
	s.Push(3000)
	fmt.Println(s.Pop())

}

/*
You have a list of URLs. Write a Go program that uses goroutines to fetch the content of each URL concurrently. 
Limit the number of simultaneous goroutines to a reasonable value (e.g., 2).
*/

type URLContents struct{
	urls []string
}

func TestULRS(t *testing.T){
	urlfetch:=&URLContents{urls: make([]string,0)}
	urlfetch.urls=append(urlfetch.urls, "https://www.google.com/")
	urlfetch.urls=append(urlfetch.urls, "https://www.yahoo.com/")
	urlfetch.urls=append(urlfetch.urls, "https://www.facebook.com/")
	urlfetch.urls=append(urlfetch.urls, "https://about.meta.com/")
	urlfetch.urls=append(urlfetch.urls, "https://www.instagram.com/")
	urlfetch.urls=append(urlfetch.urls, "https://www.google.com/maps")

	


}
