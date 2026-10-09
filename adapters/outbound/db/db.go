package db

import (
	"context"
	"errors"
	"fmt"
	"go_mindmap/core/domain"

	_ "github.com/mattn/go-sqlite3"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

type Neo4jAdapter struct {
	Driver neo4j.Driver
}

func SetupDriver(uri, username, password string) (neo4j.Driver, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {

		return nil, err
	}

	ctx := context.Background()
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to Neo4j!")
	return driver, nil
}

func (f Neo4jAdapter) MindmapToNeo(ctx context.Context, node domain.Node) error {
	session := f.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {

		query := "MERGE (t:Node {name: $name, id:$id}) RETURN t"
		params := map[string]any{"name": node.Name, "id": node.Id}

		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		return result.Collect(ctx)
	})

	return err
}

func (f Neo4jAdapter) NodesFromNeo(ctx context.Context) ([]domain.Node, error) {
	query := "MATCH (n) RETURN (n)"
	raw, err := neo4j.ExecuteQuery(ctx, f.Driver, query, map[string]any{}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		return nil, err
	}

	var nodes []domain.Node
	for _, record := range raw.Records {
		//fmt.Println(record.Values[0])

		neoNode, ok := record.Values[0].(neo4j.Node)
		if !ok {
			err := errors.New("98 Problems")
			fmt.Println("could not cast record to neo4j.Node, skipping")
			return nil, err
		}

		name, ok := neoNode.Props["name"].(string)
		if !ok {
			err := errors.New("Problem beim Type-Casting von node property zu string aufgetreten.")
			return nil, err
		}

		nodes = append(nodes, domain.Node{
			Id:   neoNode.ElementId,
			Name: name,
		})
	}

	return nodes, nil
}

func (f Neo4jAdapter) ConnectNodes(ctx context.Context, nameA string, nameB string, kind string, color string, id string) error {
	session := f.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `
		MATCH (a:Node {name: $NameA}), (b:Node {name: $NameB})
		CREATE (a)-[r:COLOR_CONNECTED {kind:$kind, color:$color, id:$id}]->(b)
		`

		params := map[string]any{"NameA": nameA, "NameB": nameB, "kind": kind, "color": color, "id": id}

		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}
		return result.Collect(ctx)

	})
	return err
}

func (f Neo4jAdapter) GetConnections(ctx context.Context) ([]domain.Relationship, error) {
	query := "MATCH (a)-[r]->(b) RETURN type(r), r"
	raw, err := neo4j.ExecuteQuery(ctx, f.Driver, query, map[string]any{}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		return nil, err
	}

	var cons []domain.Relationship
	for _, record := range raw.Records {
		//fmt.Println(record.Values[0])
		relType, _ := record.Values[0].(string)
		neoCon, ok := record.Values[1].(neo4j.Relationship)
		if !ok {
			err := errors.New("97 Problems")
			fmt.Println("could not cast record to neo4j.Relationships, skipping")
			return nil, err
		}

		kind, ok := neoCon.Props["kind"].(string)
		if !ok {
			kind = relType
		}

		color, ok := neoCon.Props["color"].(string)
		if !ok {
			color = "#888888"
			// err := errors.New("Problem beim Type-Casting von relationship property zu string aufgetreten.")
		}

		id, ok := neoCon.Props["id"].(string)
		if !ok {
			err := errors.New("Problem beim Auslesen der ID eines Relationships")
			return nil, err
		}

		cons = append(cons, domain.Relationship{
			NameA: neoCon.StartElementId,
			NameB: neoCon.EndElementId,
			Kind:  kind,
			Color: color,
			Id:    id,
		})
	}

	return cons, nil
}

func (f Neo4jAdapter) GetAll(ctx context.Context) (domain.Collection, error) {
	query := "MATCH (a) OPTIONAL MATCH (a)-[r]->(b) RETURN a, r, b"
	raw, err := neo4j.ExecuteQuery(ctx, f.Driver, query, map[string]any{}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))

	var coll domain.Collection
	coll.Nodes = make(map[string]domain.Node)
	// fmt.Println(coll)

	if err != nil {
		return coll, err
	}

	for _, record := range raw.Records {
		aRaw, exists := record.Get("a")
		if !exists {
			return coll, errors.New("got result without valid node - should not happen")
		}
		a, ok := aRaw.(neo4j.Node)
		if !ok {
			return coll, errors.New("type casting zur Node schlug fehl - should not happen")
		}
		aId, ok := a.Props["id"].(string)
		if !ok {
			return coll, errors.New("type casting Id of node failed -snh")
		}
		aName, ok := a.Props["name"].(string)
		if !ok {
			return coll, errors.New("type casting name of node failed -snh")
		}
		coll.Nodes[aId] = domain.Node{Id: aId, Name: aName}
		rRaw, exists := record.Get("r")
		if !exists || rRaw == nil {
			continue
		}
		bRaw, exists := record.Get("b")
		if !exists || bRaw == nil {
			continue
		}
		b, ok := bRaw.(neo4j.Node)
		if !ok {
			return coll, errors.New("B - type casting zur Node schlug fehl - should not happen")
		}
		bName, ok := b.Props["name"].(string)
		if !ok {
			return coll, errors.New("type casting B node name failed when casting for assigning relationship endpoint - snh")
		}
		r, ok := rRaw.(neo4j.Relationship)
		if !ok {
			return coll, errors.New("type casting zum Relationship schlug fehl - should not happen")
		}
		rKind, ok := r.Props["kind"].(string)
		if !ok {
			return coll, errors.New("kind not found in relationship - snh")
		}
		rColor, ok := r.Props["color"].(string)
		if !ok {
			return coll, errors.New("did not find relationship color - snh")
		}
		rId, ok := r.Props["id"].(string)
		if !ok {
			return coll, errors.New("id not found in relationship - snh")
		}
		coll.Relationships = append(coll.Relationships, domain.Relationship{NameA: aName, NameB: bName, Kind: rKind, Color: rColor, Id: rId})
	}
	return coll, nil
}
