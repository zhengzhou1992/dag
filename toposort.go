package dag

type TOPOAlgoType uint8

const (
	TOPO_Naive TOPOAlgoType = iota
	TOPO_RefCount
	TOPO_DFS
)

type TOPOResolver interface {
	Sort(dag *DAG) []*Node
}

type naiveResolver struct {
}

func newNaiveResolver() *naiveResolver {
	return &naiveResolver{}
}

func naiveSort(dag *DAG, toSortNodes []*Node) []*Node {
	//init to sort nodes
	if len(toSortNodes) == 0 {
		for _, node := range dag.nodes {
			toSortNodes = append(toSortNodes, node)
		}
	}

	//base case
	if len(toSortNodes) == 1 {
		return toSortNodes
	}

	//remove a node from to sort set
	v := toSortNodes[0]
	toSortNodes = toSortNodes[1:]

	//recursion
	seq := naiveSort(dag, toSortNodes)
	minI := 0
	for i, u := range seq {
		if _, exist := u.toNodes[v.Id]; exist {
			minI = i + 1
		}
	}
	seq = append(seq[:minI], append([]*Node{v}, seq[minI:]...)...)
	return seq
}

func (n *naiveResolver) Sort(dag *DAG) []*Node {
	return naiveSort(dag, nil)
}

type refCountResolver struct {
}

func newRefCountResolver() *refCountResolver {
	return &refCountResolver{}
}

func (r *refCountResolver) Sort(dag *DAG) []*Node {
	//in degree
	refCount := make(map[NodeId]uint)
	for _, node := range dag.nodes {
		refCount[node.Id] = 0
	}
	for _, node := range dag.nodes {
		for toId, _ := range node.toNodes {
			refCount[toId] += 1
		}
	}

	//init queue
	q := make([]NodeId, 0)
	res := make([]*Node, 0, len(dag.nodes))
	for toId, count := range refCount {
		if count == 0 {
			q = append(q, toId)
		}
	}

	//resolve dependency
	for len(q) != 0 {
		u := q[0]
		q = q[1:]
		res = append(res, dag.nodes[u])
		for toId, _ := range dag.nodes[u].toNodes {
			refCount[toId] -= 1
			if refCount[toId] == 0 {
				q = append(q, toId)
			}
		}
	}
	return res
}

type dfsResolver struct {
}

func newDfsResolver() *dfsResolver {
	return &dfsResolver{}
}

func (d *dfsResolver) Sort(dag *DAG) []*Node {
	seen := make(map[NodeId]bool)
	res := make([]*Node, 0)

	var dfs func(node *Node)

	dfs = func(node *Node) {
		if _, exist := seen[node.Id]; exist {
			return
		}
		seen[node.Id] = true
		for id, _ := range node.toNodes {
			dfs(dag.nodes[id])
		}
		res = append(res, node)
	}

	for _, n := range dag.nodes {
		dfs(n)
	}

	for i := len(res)/2 - 1; i >= 0; i-- {
		opp := len(res) - 1 - i
		res[i], res[opp] = res[opp], res[i]
	}
	return res
}

func TopoSortFactory(algo TOPOAlgoType) TOPOResolver {
	switch algo {
	case TOPO_Naive:
		return newNaiveResolver()
	case TOPO_RefCount:
		return newRefCountResolver()
	case TOPO_DFS:
		return newDfsResolver()
	default:
		return nil
	}
}

func TopoSort(algo TOPOAlgoType, dag *DAG) []*Node {
	sorter := TopoSortFactory(algo)
	return sorter.Sort(dag)
}
