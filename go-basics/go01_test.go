package gobasics

import
(
	"testing"
	"fmt"
)

func TestX(t *testing.T){
	fmt.Println()
}

func TestBasics(t *testing.T){
	hallo()
}
func hallo(){
	fmt.Println("hallo Welt")
	var x int =101
	var str1 string="Hallo Welt"
	fmt.Printf("%d and %s\n",x,str1)
	// if else
	variable:=100
	if variable==100{
		fmt.Printf("%v\n",variable)
	}else{
		fmt.Printf("%v\n","no match")
	}
	//loops
	for i:=0;i<3;i++{
		fmt.Printf("%d\t",i)
	}
	y:=10
	res:=func (x int)(int)  {
		fmt.Printf("%d\n",x*x)
		return x*x
	}(y)
	fmt.Printf("%d\t",res)
}
func TestArraysNSlices(t *testing.T){
	arraysnslices()
}
func arraysnslices(){
	arr:=[5]int{1,2,3,4,5}
	
	slc:=arr[0:]
	slc = append(slc, 81)
	fmt.Println(arr,slc)
}
func TestMaps(t *testing.T){
	mindmap()
}
func mindmap(){
	session:=make(map[int]string , 0)
	session[1]="one"
	session[2]="two"
	for i,s:=range session{
		fmt.Println(i,s)
	}
	fmt.Println(session)
}
