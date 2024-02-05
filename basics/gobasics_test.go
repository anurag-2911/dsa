package basics

import
(
	"fmt"
	"testing"
)

func TestTypeAssertion(t *testing.T){
	defer func(){
		if r:=recover();r!=nil{
			
		}
	}()
	fmt.Print()
	var i interface{}="hallo"
	val,ok:=i.(string)
	if ok{
		fmt.Println(val)
	}
	var k interface{}="Hallo"
	v,ok:=k.(int)
	if ok{
		fmt.Println(v)
	}


}