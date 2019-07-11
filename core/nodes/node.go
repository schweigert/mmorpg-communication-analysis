package nodes

// Node interface store all node methods
type Node interface {
	Attach(nodes ...Node)
	Children() []Node
	Init()
}

// NewBaseNode instance a new BaseNode
func NewBaseNode(children ...Node) Node {
	if children == nil {
		children = []Node{}
	}

	return &BaseNode{children: children}
}
