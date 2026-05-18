package usecases

import (
	"context"
	db "go_mindmap/adapters/outbound/db"
	"go_mindmap/core/domain"
)

type MapActions struct {
	// Prefix string
	Nodes map[string]domain.Node
	DB    db.Neo4jAdapter
}

func (f MapActions) GetNodeByID(id string) (domain.Node, bool) {
	node, exists := f.Nodes[id]
	return node, exists
	// return f.Prefix + id
}

func (f MapActions) SaveNode(node domain.Node) error {
	err := f.DB.MindmapToNeo(context.Background(), node.Content)
	if err != nil {
		return err
	}
	return nil
}

func (f MapActions) GetNodes() error {
	err := f.DB.NodesFromNeo(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (f MapActions) ConnectNodesFE(node1 domain.Node, node2 domain.Node) error {
	err := f.DB.ConnectNodes(context.Background(), node1, node2, "IS_CONNECTED_TO")
	if err != nil {
		return err
	}
	return nil
}
