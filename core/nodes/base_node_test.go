package nodes

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type BaseNodeSuite struct {
	gspec.Suite
}

func (suite *BaseNodeSuite) SetupTest() {
	suite.Suite.SetupTest()
}

func (suite *BaseNodeSuite) TestChildren() {
	children := []Node{NewBaseNode(), NewBaseNode()}
	node := NewBaseNode(children...)

	suite.Equal(children, node.Children())
}

func (suite *BaseNodeSuite) TestAttach() {
	children := []Node{NewBaseNode(), NewBaseNode()}
	node := NewBaseNode()

	suite.Equal([]Node{}, node.Children())

	node.Attach(children...)
	suite.Equal(children, node.Children())
}

func TestBaseNodeSuite(t *testing.T) {
	gspec.Run(t, new(BaseNodeSuite))
}
