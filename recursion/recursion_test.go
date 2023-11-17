package recursion

import (
	"fmt"
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
	substr:=s[1:]
	first:=s[0]
	strfirst:=string(first)
	fmt.Print(strfirst)
	fmt.Print(" ")
	return revstring(substr) + string(first)
}
