package gobasics

import (
	"fmt"
	"testing"
	"time"
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

func TestStructMethods(t *testing.T){
	bundle:=&Bundle{}
	bundle.Add("1","one")
}
type Bundle struct{
	list map[string]string
}
func NewBundle()*Bundle{
	return &Bundle{list:map[string]string{}}
}
func(b *Bundle)Add(id string,Name string){
	if b.list==nil{
		b.list=make(map[string]string)
	}
	b.list[id]=Name
}
func TestIFace(t *testing.T){
	var action Action
	action=&FileCopy{}
	fmt.Println(action.actionName(),action.processAction())
	action=&InstallDir{}
	fmt.Println(action.actionName(),action.processAction())
}

type Action interface{
	processAction()bool
	actionName()string
}

type FileCopy struct{

}
func(fc *FileCopy)processAction()bool{
	time.Sleep(2*time.Second)
	return true
}
func(fc *FileCopy)actionName()string{
	return "FileCopy"
}
type InstallDir struct{

}
func(id *InstallDir)processAction()bool{
	time.Sleep(time.Second*2)
	return true
}
func(id *InstallDir)actionName()string{
	return "InstallDir"
}