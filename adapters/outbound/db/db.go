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

// func setUpdb() (*sql.db, error) { //Roman fragen nach *xyz logik
// 	db, err := sql.Open("sqlite3", "./mindmap.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//defer db.Close()
// 	creationStatement := `
// 	CREATE TABLE IF NOT EXISTS mmaps (
// 		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
// 		node JSON
// 	);
// 	`

// 	_, err = db.Exec(creationStatement)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	log.Println("Table 'mmap' angelegt")
// 	return db, nil
// }

// func CommitMmapTodb(f *sql.db, g domain.Node) error {
// 	query := "INSERT INTO mmaps(node) VALUES(?) RETURNING id"
// 	_, err := f.Exec(query, g.Content)
// 	if err != nil {
// 		log.Printf("failure to insert node: %v", err)
// 		return err
// 	}

// 	log.Println("Node saved in db")
// 	return nil
// }

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

func (f Neo4jAdapter) MindmapToNeo(ctx context.Context, content string) error {
	session := f.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {

		query := "MERGE (t:Node {name: $name}) RETURN t"
		params := map[string]any{"name": content}

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
		fmt.Println(record.Values[0])

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

func (f Neo4jAdapter) ConnectNodes(ctx context.Context, nameA string, nameB string, kind string, color string) error {
	session := f.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `
		MATCH (a:Node {name: $NameA}), (b:Node {name: $NameB})
		CREATE (a)-[r:COLOR_CONNECTED {kind:$kind, color:$color}]->(b)
		`

		params := map[string]any{"NameA": nameA, "NameB": nameB, "kind": kind, "color": color}

		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}
		return result.Collect(ctx)
	})
	return err
}

func (f Neo4jAdapter) GetConnections(ctx context.Context) ([]domain.Relationship, error) {
	query := "MATCH (a)-[r]->(b) RETURN r"
	raw, err := neo4j.ExecuteQuery(ctx, f.Driver, query, map[string]any{}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		return nil, err
	}

	var cons []domain.Relationship
	for _, record := range raw.Records {
		fmt.Println(record.Values[0])
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
			err := errors.New("Problem beim Type-Casting von relationship property zu string aufgetreten.")
			return nil, err
		}

		cons = append(cons, domain.Relationship{
			NameA: neoCon.StartElementId,
			NameB: neoCon.EndElementId,
			Kind:  kind,
			Color: color,
		})
	}

	return cons, nil
}

// func (f Neo4jAdapter) SingleNodeFromNeo(ctx context.Context, id string) error {
// 	session := f.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
// 	defer session.Close(ctx)

// 	_, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
// 		query := "MATCH (n) RETURN (n)"
// 		result, err := tx.Run(ctx, query, nil)
// 		if err != nil {
// 			return nil, err
// 		}

// 		fmt.Println("______Los gehts, hier alle Nodes________")

// 		for result.Next(ctx) {
// 			record := result.Record()
// 			fmt.Println(record.Values...)
// 		}

// 		return nil, result.Err()

// 	})

// 	return err
// }

// query := `
// MATCH (a:Node {name: $IdA}), (b:Node{name: $IdB})
// CREATE (a)-[r:IS_CONNECTED_TO]->(b)
// `
