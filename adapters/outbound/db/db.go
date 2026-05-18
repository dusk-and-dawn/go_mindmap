package db

import (
	"context"
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

func (f Neo4jAdapter) MindmapToNeo(ctx context.Context, topicName string) error {
	session := f.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {

		query := "MERGE (t:Node {name: $name}) RETURN t"
		params := map[string]any{"name": topicName}

		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		return result.Collect(ctx)
	})

	return err
}

func (f Neo4jAdapter) NodesFromNeo(ctx context.Context) error {
	session := f.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	_, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := "MATCH (n) RETURN (n)"
		result, err := tx.Run(ctx, query, nil)
		if err != nil {
			return nil, err
		}

		fmt.Println("______Los gehts, hier alle Nodes________")

		for result.Next(ctx) {
			record := result.Record()
			fmt.Println(record.Values...)
		}

		return nil, result.Err()

	})

	return err
}

func (f Neo4jAdapter) ConnectNodes(ctx context.Context, nodeA domain.Node, nodeB domain.Node, relationshipA string) error {
	session := f.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `
		MATCH (a:Node {name: $nameA}), (b:Node{name: $nameB})
		MERGE (a)-[r:IS_CONNECTED_TO]->(b)
		Return r
		`
		params := map[string]any{"nameA": nodeA, "nameB": nodeB}

		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}
		return result.Collect(ctx)
	})
	return err
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
