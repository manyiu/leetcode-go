package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if visitedNode, ok := visited[node]; ok {
		return visitedNode
	}

	cloneNode := &Node{Val: node.Val}
	visited[node] = cloneNode

	for _, neighbor := range node.Neighbors {
		cloneNode.Neighbors = append(cloneNode.Neighbors, clone(neighbor, visited))
	}

	return cloneNode
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := map[*Node]*Node{}
	return clone(node, visited)
}
