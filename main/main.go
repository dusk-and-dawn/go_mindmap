// main.go
package main

import (
	"fmt"
	ihttp "go_mindmap/adapters/inbound/http"
	"go_mindmap/adapters/outbound/db"
	"go_mindmap/core/domain"
	"go_mindmap/core/usecases"
	"log"
	"net/http"
)

func getHandler() {

}

func main() {

	fmt.Println("In main angekommen")

	// conn, err := setUpdb()
	// if err != nil {
	// 	fmt.Println("keine Connection zur db moeglich", "error ist:", err)
	// 	return
	// }

	mux := http.NewServeMux()

	a := domain.Node{
		Id:       "01",
		Content:  "mega test node",
		Children: []string{"a", "b"},
		X:        160,
		Y:        200}

	b := domain.Node{
		Id:       "02",
		Content:  "beste Teste",
		Children: []string{"a", "b"},
		X:        160,
		Y:        200}

	c := []domain.Node{a, b}

	driver, err := db.SetupDriver("neo4j://127.0.0.1:7687", "neo4j", "Derunendliche")
	if err != nil {
		fmt.Println("kaputt")
		return
	}

	mm2neo := db.Neo4jAdapter{Driver: driver}

	if err != nil {
		fmt.Println("kaputt2")
		return
	}
	fmt.Println("weggeschrieben in Neo4j")
	// commitMmapTodb(conn, a)
	// var testStringID string
	// var testStringContent string
	// err = conn.QueryRow("SELECT * FROM mmaps WHERE id = ?", 1).Scan(&testStringID, &testStringContent)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println(testStringID, testStringContent)

	nodes := map[string]domain.Node{}

	for _, value := range c {
		nodes[value.Id] = value
	}

	mh := ihttp.MapHandler{MAusecase: usecases.MapActions{Nodes: nodes, DB: mm2neo}}

	mux.HandleFunc("GET /map/{id}", mh.GetNode) // dies in den http adapter

	mux.HandleFunc("GET /makemap", mh.ShowWaiting) // dies in den http adapter

	mux.HandleFunc("POST /makemap", mh.CreateNode) // dies in den http adapter

	log.Fatal(http.ListenAndServe(":8080", mux)) // dies in den http adapter

}
