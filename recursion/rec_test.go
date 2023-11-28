package recursion

import (
	"fmt"
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

	if(len(s)==0 || len(s)==1){
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
func TestXPlain(t *testing.T){
	fmt.Println(xpalin("racecar"))
	fmt.Println(xpalin("racegurram"))
}

func xdectobin(s string)string{
	return ""
}