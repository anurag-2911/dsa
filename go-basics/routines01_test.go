package gobasics

import
(
	"testing"
	"fmt"
)

func TestXxxY(t *testing.T) {
	fmt.Println()
}

func TestSlc(t *testing.T){
	arr:=make([]int,0,10)
	arr09:=[3]int{0,1,2}
	for i:=0;i<3;i++{
		arr = append(arr, i)
	}
	fmt.Println(arr,arr09)
	changeslc(arr,arr09)
	fmt.Println(arr,arr09)

	list:=make(map[int]string)
	list[1]="one"
	list[20]="two"
	for i,val:=range list{
		fmt.Println(i,val)
	}

	ch:=make(chan int)
	go func (ch chan int)  {
		ch<-101
		close(ch)
	}(ch)
	for v:=range ch{
		fmt.Println(v)
	}
	fmt.Println("all done")
}
func changeslc(arr2 []int,arr3 [3]int){
	arr2[2]=901
	arr3[2]=109
}

func TestIXFace(t *testing.T){
	cl:=&ConsoleLogger{}
	logme(cl,"hallo")
}
type logger interface{
	logmessage(msg string)
}

type ConsoleLogger struct{

}
func(cl *ConsoleLogger)logmessage(msg string){
	fmt.Println(msg)
}
func logme(l logger,msg string){
	l.logmessage(msg)
}

func TestPR(t *testing.T){
	pan()
	fmt.Println("all done")

}

func pan() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from a panic ", r)
		}
	}()
	causePanic()
	fmt.Println("after a panic")
}
func causePanic(){
	panic("going to panic")
}