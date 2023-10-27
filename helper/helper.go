package helper

import
(
	"fmt"
)

func PrintFormatted(data interface{}){
	fmt.Print(data)
	fmt.Print(" | ")
}