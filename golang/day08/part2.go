package main

func complicatedSum(n *node) (sum int) {
	if len(n.children) == 0 {
		for _, metadata := range n.metadata {
			sum += metadata
		}
	} else {
		for _, value := range n.metadata {
			index := value - 1
			if index < len(n.children) {
				sum += complicatedSum(n.children[index])
			}
		}
	}
	return
}
