package graphs

type GraphNode struct {
	Name     string
	Children []*GraphNode
}

func (n *GraphNode) BreadthFirstSearch(arr []string) []string {
	queue := []*GraphNode{n}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		arr = append(arr, curr.Name)
		for _, child := range curr.Children {
			queue = append(queue, child)
		}
	}
	return arr
}