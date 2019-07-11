package nodes

import (
	"sync"
	"time"
)

// BaseNode implements all node interface methods
type BaseNode struct {
	children      []Node
	childrenMutex sync.Mutex
}

// Init children and this node
func (node *BaseNode) Init() {
	for _, child := range node.Children() {
		child.Init()
	}
}

// Update children and this node
func (node *BaseNode) Update(deltaTime time.Duration) {
	for _, child := range node.Children() {
		child.Update(deltaTime)
	}
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

// Disinherit an node from children
func (node *BaseNode) Disinherit(child Node) {
	node.childrenMutex.Lock()
	defer node.childrenMutex.Unlock()

	index := -1

	for id, searchChild := range node.children {
		if searchChild == child {
			index = id
		}
	}

	if index != -1 {
		node.children = append(node.children[:index], node.children[index+1:]...)
	}
}
