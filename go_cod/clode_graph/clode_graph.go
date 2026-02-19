package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	copies := make([]*Node, 101)

	dfs(node, copies)

	return copies[node.Val]
}

func dfs(node *Node, copies []*Node) {
	newNode := new(Node)
	newNode.Val = node.Val

	copies[node.Val] = newNode

	for _, neighbor := range node.Neighbors {
		if copies[neighbor.Val] == nil {
			dfs(neighbor, copies)
		}

		newNode.Neighbors = append(newNode.Neighbors, copies[neighbor.Val])
	}
}

func main() {
	// Example usage:
	// Create a graph with 4 nodes
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	// Clone the graph
	clonedGraph := cloneGraph(node1)

	// Print the cloned graph's nodes and their neighbors
	println("Cloned Graph:")
	for _, node := range []*Node{clonedGraph, clonedGraph.Neighbors[0], clonedGraph.Neighbors[0].Neighbors[0], clonedGraph.Neighbors[0].Neighbors[1]} {
		println("Node:", node.Val)
		print("Neighbors: ")
		for _, neighbor := range node.Neighbors {
			print(neighbor.Val, " ")
		}
		println()
	}
}
