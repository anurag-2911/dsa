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

type YBinaryTree struct {
	Root *YNode
}
type YNode struct {
	Data  int
	Left  *YNode
	Right *YNode
}

func (tree *YBinaryTree) AddNode(data int) {
	newNode := &YNode{Data: data}
	if tree.Root == nil {
		tree.Root = newNode
		return
	}
	AddNodes(tree.Root, newNode)
}

func AddNodes(current *YNode, newNode *YNode) {
	if newNode.Data < current.Data {
		//left
		if current.Left == nil {
			current.Left = newNode
			return
		}
		AddNodes(current.Left,newNode)
	} else {
		//right
		if current.Right == nil {
			current.Right = newNode
			return
		}
		AddNodes(current.Right,newNode)
	}
}

func TestYBinaryTree(t *testing.T) {
	fmt.Println("binary tree operations")
	data := []int{5, 3, 7, 1, 4, 6,8}
	tree := &YBinaryTree{}
	for _, v := range data {
		tree.AddNode(v)
	}

}
func TestYBinaryTreeInOrder(t *testing.T){
	fmt.Println("binary tree operations")
	data := []int{5, 3, 7, 1, 4, 6,8}
	tree := &YBinaryTree{}
	for _, v := range data {
		tree.AddNode(v)
	}
	tree.InOrder()
	tree.PreOrder()

}
func(tree *YBinaryTree)InOrder(){
	YInOrderTraversal(tree.Root)
}
func YInOrderTraversal(root *YNode){
	if root==nil{
		return
	}
	YInOrderTraversal(root.Left)// visit the left subtree
	fmt.Print(root.Data) // visit the node
	fmt.Print(" > ")
	YInOrderTraversal(root.Right)// visit the right subtree
}
func (tree *YBinaryTree)PreOrder(){
	YPreOrder(tree.Root)
}
func YPreOrder(root *YNode){
	if root==nil{
		return
	}
	fmt.Print(root.Data)
	fmt.Print(" > ")
	YPreOrder(root.Left)
	YPreOrder(root.Right)
}
func TestYBinaryPreOrder(t *testing.T){
	fmt.Println("binary tree operations")
	data := []int{5, 3, 7, 1, 4, 6,8}
	tree := &YBinaryTree{}
	for _, v := range data {
		tree.AddNode(v)
	}
	
	tree.PreOrder()

}
func (tree *YBinaryTree)YPostOrder(){
	YPostOrderTraversal(tree.Root)
}
func YPostOrderTraversal(root *YNode){
	if root==nil{
		return
	}
	YPostOrderTraversal(root.Left)
	YPostOrderTraversal(root.Right)
	fmt.Print(root.Data)
	fmt.Print(" > ")
}
func TestYBinaryPostOrder(t *testing.T){
	fmt.Println("binary tree operations")
	data := []int{5, 3, 7, 1, 4, 6,8}
	tree := &YBinaryTree{}
	for _, v := range data {
		tree.AddNode(v)
	}
	
	tree.YPostOrder()

}