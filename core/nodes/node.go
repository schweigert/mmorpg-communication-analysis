package nodes

import "time"

// Node interface
type Node interface {
	Children() []Node
	Update(deltaTime time.Duration)
	Attach(newChildren ...Node)
}

// NewHierarchicalNode create a tree based node
func NewHierarchicalNode() Node {
	return &HierarchicalNode{children: []Node{}}
}
