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
	err := f.DB.MindmapToNeo(context.Background(), node.Name)
	if err != nil {
		return err
	}
	return nil
}

func (f MapActions) GetNodes() ([]domain.Node, error) {
	nodes, err := f.DB.NodesFromNeo(context.Background())
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (f MapActions) ConnectNodesFE(nameA string, nameB string, kind string, color string) error {
	err := f.DB.ConnectNodes(context.Background(), nameA, nameB, kind, color)
	if err != nil {
		return err
	}
	return nil
}

func (f MapActions) GetConnections() ([]domain.Relationship, error) {
	connections, err := f.DB.GetConnections(context.Background())
	if err != nil {
		return nil, err
	}
	return connections, nil
}
