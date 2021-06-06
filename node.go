package dag

import "github.com/google/uuid"

type NodeId uuid.UUID
type EdgeWeight int

type Node struct {
	Id      NodeId
	data    interface{}
	toNodes map[NodeId]EdgeWeight
}

func NewNode(data interface{}) (*Node, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &Node{
		Id:      NodeId(id),
		data:    data,
		toNodes: make(map[NodeId]EdgeWeight),
	}, nil
}

func (n *Node) AddEdge(to *Node) {
	n.toNodes[to.Id] = 1 // TODO support weight and other edge propeties
}
