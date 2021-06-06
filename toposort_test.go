package dag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type TOPOSortTestSuit struct {
	suite.Suite
	demoDAG *DAG
}

func (suite *TOPOSortTestSuit) SetupTest() {
	dag := NewDAG()
	na, _ := NewNode("a")
	nb, _ := NewNode("b")
	nc, _ := NewNode("c")
	nd, _ := NewNode("d")
	ne, _ := NewNode("e")
	nf, _ := NewNode("f")

	_ = dag.AddNode(nf)
	_ = dag.AddNode(ne)
	_ = dag.AddNode(nd)
	_ = dag.AddNode(nc)
	_ = dag.AddNode(na)
	_ = dag.AddNode(nb)

	dag.AddEdge(na, nb)
	dag.AddEdge(na, nf)
	dag.AddEdge(nb, nf)
	dag.AddEdge(nb, nc)
	dag.AddEdge(nb, nd)
	dag.AddEdge(nc, nd)
	dag.AddEdge(nd, nf)
	dag.AddEdge(nd, ne)
	dag.AddEdge(ne, nf)

	suite.demoDAG = dag
}

func (t *TOPOSortTestSuit) verifySortResult(res []*Node) {
	expectSeq := []string{"a", "b", "c", "d", "e", "f"}
	assert.Equal(t.T(), len(expectSeq), len(res))

	for i, node := range res {
		assert.Equal(t.T(), node.data.(string), expectSeq[i])
	}
}

func (t *TOPOSortTestSuit) TestRefCountTOPOSortAlgo() {
	res := TopoSort(TOPO_RefCount, t.demoDAG)
	t.verifySortResult(res)
}

func (t *TOPOSortTestSuit) TestNaiveTOPOSortAlgo() {
	res := TopoSort(TOPO_Naive, t.demoDAG)
	for i, node := range res {
		fmt.Printf("idx %d -,  node value - %s\n", i, node.data)
	}
	t.verifySortResult(res)
}

func (t *TOPOSortTestSuit) TestDFSTOPOSortAlgo() {
	res := TopoSort(TOPO_DFS, t.demoDAG)
	t.verifySortResult(res)
}

func TestTOPOSortAlgos(t *testing.T) {
	suite.Run(t, new(TOPOSortTestSuit))
}
