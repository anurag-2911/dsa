package util

import (
	"reflect"
	"testing"
)

func TestMergeArrays(t *testing.T){
	got:=merge([]int{2, 5, 6, 9, 11},[]int{1, 1, 3, 5, 8})
	want:=[]int{1, 1, 2, 3, 5, 5, 6, 8, 9, 11}
	if !reflect.DeepEqual(got,want){
		t.Errorf("got %v,want %v",got,want)
	}
}
func TestMergeSimilarArrays(t *testing.T){
	got:=merge([]int{1, 1, 1, 9, 11},[]int{1, 1, 1, 1, 8})
	want:=[]int{1, 1, 1, 1, 1, 1, 1, 8, 9, 11}
	if !reflect.DeepEqual(got,want){
		t.Errorf("got %v,want %v",got,want)
	}
}
func TestMergeNegativeAndZeros(t *testing.T){
	got:=merge([]int{-10,-1,0,0},[]int{0, 0})
	want:=[]int{-10,-1,0,0,0,0}
	if !reflect.DeepEqual(got,want){
		t.Errorf("got %v,want %v",got,want)
	}
}

func TestMergeOneEmptyArray(t *testing.T){
	got:=merge([]int{},[]int{1,2,3})
	want:=[]int{1,2,3}
	if !reflect.DeepEqual(got,want){
		t.Errorf("got %v,want %v",got,want)
	}
}
