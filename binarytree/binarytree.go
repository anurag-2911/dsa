package binarytree

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (tree *BinaryTree) Insert(data int) {
	node := &Node{Data: data}
	if tree.Root == nil {
		tree.Root = node
	} else {
		insertData(tree.Root, node)
	}
}
func insertData(node *Node, newNode *Node) {
	if node.Data > newNode.Data {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertData(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertData(node.Right, newNode)
		}
	}

}
func (tree *BinaryTree) InOrderTraversal(root *Node, f func(int)) {
	inOrder(root, f)
}
func inOrder(node *Node, f func(int)) {
	if(node==nil){
		return
	}
	inOrder(node.Left,f)
	f(node.Data)
	inOrder(node.Right,f)
}
