package basics

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"os"
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

	maxConcurrent:=2

	var wg sync.WaitGroup

	semaphore:=make(chan struct{},maxConcurrent)

	for _,url:=range urlfetch.urls{
		wg.Add(1)
		go func (url string)  {
			defer wg.Done()
			semaphore<-struct{}{}

			//http request
			resp,err:=http.Get(url)
			if err!=nil{
				fmt.Printf("error fetching %s url %v\n",url,err)
				return
			}
			defer resp.Body.Close()
			content,err:=io.ReadAll(resp.Body)
			if err!=nil{
				fmt.Printf("error in reading the content for url %s\n %v ",url,err)
				return
			}
			fmt.Println("content for url",url,string(content))
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

func TestFileOps(t *testing.T){
	fmt.Println(fileops("./basics_test.go"))
}
func fileops(filePath string)(string,error){
	fs,err:=os.Open(filePath)
	if err!=nil{
		fmt.Println(err)
		return "",err
	}
	defer fs.Close()
	data,err:=io.ReadAll(fs)
	if err!=nil{
		return "",err
	}
	return string(data),nil
}




