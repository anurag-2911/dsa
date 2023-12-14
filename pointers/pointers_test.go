package pointers

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	fmt.Println()
}

func TestPointerDec(t *testing.T) {
	var a int = 10
	var p *int = &a
	fmt.Println("value of a ", a)
	fmt.Println("address of a ", p)
	fmt.Println("value at address p ", *p)
}

func TestChangeValue(t *testing.T) {
	var a int = 100
	var aptr *int = &a

	fmt.Println(a)

	*aptr = 400
	fmt.Println(a)

}

func TestPassingPtrToFunction(t *testing.T) {
	var x int = 100
	increment(&x)
	fmt.Println(x)

}

func increment(x *int) {
	*x += 1
}

func TestPtrArray(t *testing.T){
	arr:=[3]int{100,200,300}
	p:=&arr

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println((*p)[0])

	(*p)[1]=900
	fmt.Println(arr)


}

func TestPtrSlice(t *testing.T){
	slice:=[]int{200,400,600}

	ptr:=&slice

	(*ptr)[2]=20

	fmt.Println(slice)

	(*ptr)[10]=2
	fmt.Println(slice)
}

func TestPtrStruct(t *testing.T){
	p:=&PX{X:100,Y:200}
	fmt.Println(p)
	fmt.Println(*p)
}
type PX struct{
	X int
	Y int
}

func TestPtrPtr(t *testing.T){
	var a int =1001
	var ptr1 *int =&a
	var ptr2 **int=&ptr1
	fmt.Println("value of a ",a)
	fmt.Println("value by ptr1 ",*ptr1)
	fmt.Println("value by ptr2",**ptr2)
	var ptr3 ***int=&ptr2
	fmt.Println(***ptr3)
}

func TestPtrToFunc(t *testing.T){
	ptr:=add(12,13)
	fmt.Println(ptr)
}

func add(x,y int)int{
	return x+y
}