package binarytree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {

	fmt.Println("binary tree operations")
	data := []int{5, 3, 7, 1, 4, 6}
	tree := &BinaryTree{}
	for _, v := range data {
		tree.Insert(v)
	}
	
}
