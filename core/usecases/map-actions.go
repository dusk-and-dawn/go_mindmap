package usecases

import (
	"context"
	db "go_mindmap/adapters/outbound/db"
	"go_mindmap/core/domain"
)

// Ist die Definition des Struct typen hier sinnvoll oder sollte ich diese auslagern?
// type Node struct {
// 	Id       string
// 	Content  map[string]any
// 	//Children []string
// 	//X        float64
// 	//Y        float64
// }

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

func (f MapActions) GetNodes() []any {
	x := []any{"test", "0"}
	return x
}
