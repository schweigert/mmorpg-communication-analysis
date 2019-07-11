package nodes

import (
	"testing"
	"time"

	"github.com/krakenlab/gspec"
)

type BaseNodeSuite struct {
	gspec.Suite
}

func (suite *BaseNodeSuite) SetupTest() {
	suite.Suite.SetupTest()
}

func (suite *BaseNodeSuite) TestInit() {
	children := []Node{NewBaseNode(), NewBaseNode()}
	node := NewBaseNode(children...)

	suite.NotPanics(func() {
		node.Init()
	})
}

func (suite *BaseNodeSuite) TestUpdate() {
	children := []Node{NewBaseNode(), NewBaseNode()}
	node := NewBaseNode(children...)

	suite.NotPanics(func() {
		node.Update(time.Second)
	})
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

func (suite *BaseNodeSuite) TestDisinherit() {
	firstChild := NewBaseNode()
	secondChild := NewBaseNode(NewBaseNode())

	suite.NotEqual(firstChild, secondChild)

	children := []Node{firstChild, secondChild}
	node := NewBaseNode(children...)

	suite.Equal(children, node.Children())

	suite.NotPanics(func() {
		node.Disinherit(firstChild)
	})

	suite.Equal([]Node{secondChild}, node.Children())
}

func TestBaseNodeSuite(t *testing.T) {
	gspec.Run(t, new(BaseNodeSuite))
}
