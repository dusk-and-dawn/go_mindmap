package usecases

import (
	"context"
	db "go_mindmap/adapters/outbound/db"
	"go_mindmap/core/domain"

	"github.com/google/uuid"
)

type MapActions struct {
	// Prefix string
	Nodes map[string]domain.Node
	DB    db.Neo4jAdapter
}

func createID() string {
	id := uuid.New().String()
	return id
}

func (f MapActions) GetNodeByID(id string) (domain.Node, bool) {
	node, exists := f.Nodes[id]
	return node, exists
	// return f.Prefix + id
}

func (f MapActions) SaveNode(node domain.Node, mindmap domain.Mindmap) error {
	node.Id = createID()
	err := f.DB.NodeToNeo(context.Background(), node, mindmap)
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

func (f MapActions) ConnectNodes(IdA string, IdB string, kind string, color string) error {
	id := createID()
	err := f.DB.ConnectNodes(context.Background(), IdA, IdB, kind, color, id)
	if err != nil {
		return err
	}
	return nil
}

func (f MapActions) GetAll(mindmapId string) (domain.Collection, error) {
	return f.DB.GetAll(context.Background(), mindmapId)
}

func (f MapActions) GetMindmaps() {

}
