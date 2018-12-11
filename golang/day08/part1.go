package main

func metadataSum(data []int) (rootNode *node, sum int) {
	currentNode := &node{
		parent:        nil,
		totalNodes:    data[0],
		totalMetadata: data[1],
	}
	rootNode = currentNode
	for i := 2; i < len(data); {
		if hasAllChildren(currentNode) && !hasAllMetadata(currentNode) {
			for !hasAllMetadata(currentNode) {
				currentNode.metadata = append(currentNode.metadata, data[i])
				i++
			}
			currentNode = currentNode.parent
		} else if !hasAllChildren(currentNode) {
			newNode := &node{
				parent:        currentNode,
				totalNodes:    data[i],
				totalMetadata: data[i+1],
			}
			currentNode.children = append(currentNode.children, newNode)
			currentNode = newNode
			i = i + 2
		}
	}
	sum = sumMetadata(rootNode)
	return
}

type node struct {
	parent        *node
	totalNodes    int
	totalMetadata int
	children      []*node
	metadata      []int
}

func hasAllChildren(n *node) bool {
	return n.totalNodes == len(n.children)
}

func hasAllMetadata(n *node) bool {
	return n.totalMetadata == len(n.metadata)
}

func sumMetadata(n *node) (sum int) {
	for _, metadata := range n.metadata {
		sum += metadata
	}
	for _, child := range n.children {
		sum += sumMetadata(child)
	}
	return
}
