DAG implements the directed acyclic graph data structure and related algorithms that I use daily, including topological sorting, the shortest path and etc.

## Usage

```go
demoDAG := NewDAG()

// create nodes
nf, _ := NewNode("f")
ne, _ := NewNode("e")
na, _ := NewNode("a")
nb, _ := NewNode("b")
nc, _ := NewNode("c")
nd, _ := NewNode("d")

// add nodes to DAG
_ = demoDAG.AddNode(nf)
_ = demoDAG.AddNode(ne)
_ = demoDAG.AddNode(nd)
_ = demoDAG.AddNode(nc)
_ = demoDAG.AddNode(na)
_ = demoDAG.AddNode(nb)

// connect nodes with edges
demoDAG.AddEdge(na, nb)
demoDAG.AddEdge(na, nf)
demoDAG.AddEdge(nb, nf)
demoDAG.AddEdge(nb, nc)
demoDAG.AddEdge(nb, nd)
demoDAG.AddEdge(nc, nd)
demoDAG.AddEdge(nd, nf)
demoDAG.AddEdge(nd, ne)
demoDAG.AddEdge(ne, nf)

// resolve node dependency
resolvedNodes := TopoSort(TOPO_RefCount, demoDAG)

for i, node := range resolvedNodes {
    fmt.Printf("idx %d -,  node value - %s\n", i, node.data)
}

//output:
//idx 0 -,  node value - a
//idx 1 -,  node value - b
//idx 2 -,  node value - c
//idx 3 -,  node value - d
//idx 4 -,  node value - e
//idx 5 -,  node value - f
```