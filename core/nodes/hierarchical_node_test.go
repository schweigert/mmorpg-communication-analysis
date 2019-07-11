package nodes

import (
	"testing"
	"time"

	"github.com/krakenlab/gspec"
)

type HierarchicalNodeSuite struct {
	gspec.Suite
}

func (suite *HierarchicalNodeSuite) SetupTest() {
	suite.Suite.SetupTest()
}

func (suite *HierarchicalNodeSuite) TestAttach() {
	node := NewHierarchicalNode()
	children := []Node{NewHierarchicalNode(), NewHierarchicalNode(), NewHierarchicalNode()}

	suite.NotPanics(func() {
		node.Attach(children...)
	})

	suite.Equal(children, node.Children())
}

func (suite *HierarchicalNodeSuite) TestUpdate() {
	node := NewHierarchicalNode()
	children := []Node{NewHierarchicalNode(), NewHierarchicalNode(), NewHierarchicalNode()}

	node.Attach(children...)

	suite.NotPanics(func() {
		node.Update(time.Millisecond * 250)
	})
}

func (suite *HierarchicalNodeSuite) TestChildren() {
	node := NewHierarchicalNode()
	children := []Node{NewHierarchicalNode(), NewHierarchicalNode(), NewHierarchicalNode()}

	suite.Equal([]Node{}, node.Children())

	node.Attach(children...)

	suite.Equal(children, node.Children())
}

func TestHierarchicalNodeSuite(t *testing.T) {
	gspec.Run(t, new(HierarchicalNodeSuite))
}
