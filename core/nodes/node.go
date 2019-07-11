package nodes

import "time"

// Node interface store all node methods
type Node interface {
	Init()
	Update(deltaTime time.Duration)

	Attach(nodes ...Node)
	Disinherit(node Node)
	Children() []Node
}

// NewBaseNode instance a new BaseNode
func NewBaseNode(children ...Node) Node {
	if children == nil {
		children = []Node{}
	}

	return &BaseNode{children: children}
}
