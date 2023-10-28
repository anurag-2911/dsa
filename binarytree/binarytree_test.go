package binarytree

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {

}
//Single Node in the tree
type Node struct{
	data int
	left *Node
	right *Node
}
//Entire Tree
type BinaryTree struct{
	root *Node
}
/*
For simplicity, create a binary search tree (BST), 
where the left child has a value less than its parent and the right child has a value greater than its parent
*/
func(this *BinaryTree)Insert(value int){
	if(this.root==nil){
		this.root=&Node{data:value}
		return
	}
	


}
