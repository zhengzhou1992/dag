package dag

type DAG struct {
	nodes map[NodeId]*Node
}

func NewDAG() *DAG {
	return &DAG{nodes: make(map[NodeId]*Node)}
}

func (d *DAG) AddNode(node *Node) error {
	if d.Exist(node) {
		return NodeExist
	}
	d.nodes[node.Id] = node
	return nil
}

func (d *DAG) Exist(node *Node) bool {
	_, exist := d.nodes[node.Id]
	return exist
}

func (d *DAG) AddEdge(from, to *Node) {
	if !d.Exist(from) {
		_ = d.AddNode(from)
	}
	if !d.Exist(to) {
		_ = d.AddNode(to)
	}
	from.AddEdge(to)
}
