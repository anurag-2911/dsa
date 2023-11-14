package graph

import "fmt"

type Node struct {
	Data int
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

func GraphMain() {
	g := &Graph{nodes: make([]*Node, 0), edges: make(map[Node][]*Node)}
	node1 := &Node{Data: 100}
	node2 := &Node{Data: 200}
	node3 := &Node{Data: 300}
	node4 := &Node{Data: 400}
	g.AddNode(node1)
	g.AddNode(node2)
	g.AddNode(node3)
	g.AddNode(node4)
	g.AddEdges(node1, node2)
	g.AddEdges(node2, node3)
	g.AddEdges(node3, node4)
	g.AddEdges(node4, node1)

	g.Show()
}

func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
	g.edges[*n] = []*Node{}
}
func (g *Graph) AddEdges(n1 *Node, n2 *Node) {
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}
func (g *Graph) Show() {
	for _, node := range g.nodes {
		fmt.Printf("%v ->", node.Data)
		neighbors := g.edges[*node]
		for _, neighbor := range neighbors {
			fmt.Printf("%v ->", neighbor.Data)
		}
		fmt.Println()
	}
}
