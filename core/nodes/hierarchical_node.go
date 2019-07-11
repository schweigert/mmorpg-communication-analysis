package nodes

import (
	"sync"
	"time"
)

// HierarchicalNode update all
type HierarchicalNode struct {
	children      []Node
	childrenMutex sync.Mutex
}

// Children of this node
func (node *HierarchicalNode) Children() []Node {
	node.childrenMutex.Lock()
	defer node.childrenMutex.Unlock()

	return node.children
}

// Update this node and your children based on deltaTime
func (node *HierarchicalNode) Update(deltaTime time.Duration) {
	for _, child := range node.Children() {
		child.Update(deltaTime)
	}
}

// Attach a new node as child
func (node *HierarchicalNode) Attach(newChildren ...Node) {
	node.childrenMutex.Lock()
	defer node.childrenMutex.Unlock()

	node.children = append(node.children, newChildren...)
}
