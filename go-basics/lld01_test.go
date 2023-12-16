package gobasics

import (
	"fmt"
	"sync"
	"testing"
)

func TestXxx(t *testing.T) {
	fmt.Println()
}

// singleton pattern

type Singleton struct{}

var
(
	instance *Singleton
	once sync.Once
)

func getInstance()*Singleton{
	once.Do(func(){
		instance=&Singleton{}
	})
	return instance
}

func TestSingleTon(t *testing.T){
	s1:=getInstance()
	s2:=getInstance()
	fmt.Println(s1==s2)
}