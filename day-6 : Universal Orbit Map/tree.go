package main

// Node to represent each node
type Node struct {
	key        string
	parent     *Node
	childIndex int
	children   []*Node
}

// OrbitTree to represent the orbit tree for traversal
type OrbitTree struct {
	rootKey string
	nodeMap map[string]*Node
}

func newOrbitTree(rootKey string) OrbitTree {
	return OrbitTree{
		rootKey: rootKey,
		nodeMap: map[string]*Node{},
	}
}

func (tree OrbitTree) addBranch(parentKey string, childKey string) {
	node := tree._addNode(childKey)
	parentNode := tree._addNode(parentKey)
	node.parent = parentNode

	parentNode.children = append(parentNode.children, node)
}

func (tree OrbitTree) _addNode(nodeKey string) *Node {
	if _, ok := tree.nodeMap[nodeKey]; !ok {
		tree.nodeMap[nodeKey] = &Node{
			key: nodeKey,
		}
	}

	return tree.nodeMap[nodeKey]
}

func (tree OrbitTree) _getNode(nodeKey string) *Node {
	if _, ok := tree.nodeMap[nodeKey]; ok {
		return tree.nodeMap[nodeKey]
	}

	return nil
}

func (tree OrbitTree) computeIndirectOrbitsCount() int {
	rootNode := tree._getNode(tree.rootKey)
	return tree._traverseAndComputeIndirectOrbitsCount(rootNode, map[*Node]int{}, 1, 0)
}

func (tree OrbitTree) _traverseAndComputeIndirectOrbitsCount(node *Node, visitedChildren map[*Node]int, currentPathNodesCount int, indirectOrbitsCount int) int {
	if _, ok := visitedChildren[node]; !ok {
		visitedChildren[node] = -1
		indirectOrbitsCount += computeIndirectOrbitsCountForBranchRoot(currentPathNodesCount)
	}

	if len(node.children) == 0 || visitedChildren[node] == len(node.children)-1 {
		if node.parent == nil {
			return indirectOrbitsCount
		}

		currentPathNodesCount--
		return tree._traverseAndComputeIndirectOrbitsCount(node.parent, visitedChildren, currentPathNodesCount, indirectOrbitsCount)
	}

	visitedChildren[node]++
	nextNode := node.children[visitedChildren[node]]
	currentPathNodesCount++

	return tree._traverseAndComputeIndirectOrbitsCount(nextNode, visitedChildren, currentPathNodesCount, indirectOrbitsCount)
}

func (tree OrbitTree) getPathFromNodeToRoot(nodeKey string) []string {
	node := tree._getNode(nodeKey)

	return tree._traversePathToRoot(node, []string{})
}

// oops not using DFS here, since have the node map and parent based linked list
func (tree OrbitTree) _traversePathToRoot(currentNode *Node, path []string) []string {
	path = append(path, currentNode.key)

	if currentNode.parent == nil {
		return path
	}

	return tree._traversePathToRoot(currentNode.parent, path)
}

func computeIndirectOrbitsCountForBranchRoot(branchNodesCount int) int {
	if branchNodesCount < 3 {
		return 0
	}

	// exclude the root node, and immediate parent node to compute indirect orbits count
	return branchNodesCount - 2
}
