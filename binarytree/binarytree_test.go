package binarytree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {

	tree := &BinaryTree{}
	populatebst(tree)
	printInOrderTraversal(tree)

}

func populatebst(tree *BinaryTree) {
	values := []int{50, 30, 70, 10, 40, 60, 80}
	for _, v := range values {
		tree.Insert(v)
	}
}

func printInOrderTraversal(tree *BinaryTree) {
	fmt.Println("in-order traversal")
	tree.InOrderTraversal(tree.root, func(val int) {
		fmt.Print(val)
		fmt.Print(" | ")
	})
}

// Single Node in the tree
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Entire Tree
type BinaryTree struct {
	root *Node
}

/*
For simplicity, create a binary search tree (BST),
where the left child has a value less than its parent and the
right child has a value greater than its parent
*/
func (this *BinaryTree) Insert(value int) {
	if this.root == nil {
		this.root = &Node{data: value}
		return
	}
	insertNode(this.root, value)

}

func insertNode(node *Node, val int) {
	if val < node.data {
		if node.left == nil {
			node.left = &Node{data: val}
		} else {
			insertNode(node.left, val)
		}
	} else {
		if node.right == nil {
			node.right = &Node{data: val}
		} else {
			insertNode(node.right, val)
		}
	}
}

// In-order Traversal
func (this *BinaryTree) InOrderTraversal(node *Node, visit func(int)) {
	if node != nil {
		this.InOrderTraversal(node.left, visit)

		visit(node.data)

		this.InOrderTraversal(node.right, visit)
	}
}

//post-order Traversal
func(this *BinaryTree)PostOrderTraversal(node *Node,visit func(int)){
	if(node!=nil){
		this.PostOrderTraversal(node.left,visit)
		this.PostOrderTraversal(node.right,visit)
		visit(node.data)
	}
}

//pre-order Traversal
func(this *BinaryTree)PreOrderTraversal(node *Node,visit func(int)){
	if(node!=nil){
		visit(node.data)
		this.PreOrderTraversal(node.left,visit)
		this.PreOrderTraversal(node.right,visit)
	}
}
