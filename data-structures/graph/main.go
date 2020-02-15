package graph

import (
	"fmt"

	"github.com/rayhaanbhikha/algorithms-data-structures-go/data-structures/node"
)

// Graph ...
type Graph struct {
	OriginNode *node.SNode
	Nodes      []*node.SNode
	Edges      map[*node.SNode][]*node.SNode
}

// NewGraph ... creates graph with origin node
func NewGraph() *Graph {

	originNode := &node.SNode{"O"}
	return &Graph{
		OriginNode: originNode,
		Nodes:      []*node.SNode{originNode},
		Edges:      make(map[*node.SNode][]*node.SNode)}
}

// AddNode ...
func (g *Graph) AddNode(newNode *node.SNode) {
	g.Nodes = append(g.Nodes, newNode)
}

// AddEdge ...
func (g *Graph) AddEdge(n1, n2 *node.SNode) {
	g.Edges[n1] = append(g.Edges[n1], n2)
}

func (g *Graph) String() string {
	str := ""
	for node, connectedNodes := range g.Edges {
		str += fmt.Sprintf("node: %s\n", node.Value)
		for _, connectedNode := range connectedNodes {
			str += connectedNode.Value + ", "
		}
		str += "\n"
	}
	return str
}
