package main

import (
	"fmt"

	"github.com/rayhaanbhikha/algorithms-data-structures-go/data-structures/graph"
	"github.com/rayhaanbhikha/algorithms-data-structures-go/data-structures/node"
	"github.com/rayhaanbhikha/algorithms-data-structures-go/data-structures/queue"
)

func main() {
	g := graph.NewGraph()
	initGraph(g)
	fmt.Println(g)
	traverseGraph(g)

}

func traverseGraph(g *graph.Graph) {
	queue := queue.Queue{}
	nodesVisited := make(map[*node.SNode]bool)

	// queue origin nodes.
	for _, node := range g.Edges[g.OriginNode] {
		queue.Enqueue(node)
	}

	for !queue.IsEmpty() {
		currentNode := queue.Dequeue()
		if nodesVisited[currentNode] {
			continue
		}

		fmt.Println(currentNode.Value)
		// get currentNodes neighbours
		neighbourNodes := g.Edges[currentNode]
		for _, node := range neighbourNodes {
			queue.Enqueue(node)
		}
		nodesVisited[currentNode] = true
	}
}

func initGraph(g *graph.Graph) {
	nodeA := &node.SNode{"A"}
	nodeB := &node.SNode{"B"}
	nodeC := &node.SNode{"C"}
	nodeD := &node.SNode{"D"}
	nodeE := &node.SNode{"E"}
	nodeF := &node.SNode{"F"}
	nodeG := &node.SNode{"G"}
	nodeH := &node.SNode{"H"}
	nodeI := &node.SNode{"I"}
	nodeJ := &node.SNode{"J"}
	nodeK := &node.SNode{"K"}
	nodeL := &node.SNode{"L"}
	nodeM := &node.SNode{"M"}
	g.AddNode(nodeA)
	g.AddNode(nodeB)
	g.AddNode(nodeC)
	g.AddNode(nodeD)
	g.AddNode(nodeE)
	g.AddNode(nodeF)
	g.AddNode(nodeG)
	g.AddNode(nodeH)
	g.AddNode(nodeI)
	g.AddNode(nodeJ)
	g.AddNode(nodeK)
	g.AddNode(nodeL)
	g.AddNode(nodeM)

	g.AddEdge(g.OriginNode, nodeA)
	g.AddEdge(g.OriginNode, nodeB)
	g.AddEdge(g.OriginNode, nodeC)
	g.AddEdge(g.OriginNode, nodeD)
	g.AddEdge(g.OriginNode, nodeJ)

	g.AddEdge(nodeA, nodeE)
	g.AddEdge(nodeA, nodeF)

	g.AddEdge(nodeB, nodeF)

	g.AddEdge(nodeC, nodeH)
	g.AddEdge(nodeC, nodeG)
	g.AddEdge(nodeC, nodeD)

	g.AddEdge(nodeD, nodeI)

	g.AddEdge(nodeJ, nodeL)
	g.AddEdge(nodeJ, nodeK)

	g.AddEdge(nodeK, nodeM)
}
