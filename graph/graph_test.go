package graph

import (
	"testing"
)

func TestXxx(t *testing.T) {

}

type Node struct {
	Value int
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}
