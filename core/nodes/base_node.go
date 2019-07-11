package nodes

import "sync"

// BaseNode implements all node interface methods
type BaseNode struct {
	children      []Node
	childrenMutex sync.Mutex
}

// Children of this node
func (node *BaseNode) Children() []Node {
	node.childrenMutex.Lock()
	defer node.childrenMutex.Unlock()

	return node.children
}

// Attach nodes in this node
func (node *BaseNode) Attach(nodes ...Node) {
	node.childrenMutex.Lock()
	defer node.childrenMutex.Unlock()

	node.children = append(node.children, nodes...)
}
