// main.go
package main

import (
	// "context"
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

	mux := http.NewServeMux()

	driver, err := db.SetupDriver("neo4j://127.0.0.1:7687", "neo4j", "Derunendliche")
	if err != nil {
		fmt.Println("Problem beim Erstellen des Drivers für Neo4j")
		return
	}

	mm2neo := db.Neo4jAdapter{Driver: driver}

	nodes := map[string]domain.Node{}

	mh := ihttp.MapHandler{MAusecase: usecases.MapActions{Nodes: nodes, DB: mm2neo}}

	mux.HandleFunc("GET /node/{id}", mh.GetNodeById) // dies in den http adapter

	mux.HandleFunc("GET /makemap", mh.ShowWaiting) // dies in den http adapter

	mux.HandleFunc("POST /node", mh.CreateNode) // dies in den http adapter

	mux.HandleFunc("GET /node", mh.GetNodes)

	mux.HandleFunc("GET /mindmap", mh.GetMindmaps)

	mux.HandleFunc("GET /mindmap/{id}", mh.GetMindmapById)

	mux.HandleFunc("POST /connect", mh.ConnectNodes)

	mux.Handle("GET /", http.FileServer(http.Dir("./frontend")))

	// mux.HandleFunc("GET /connections", mh.GetConnections)

	log.Fatal(http.ListenAndServe(":8080", mux)) // dies in den http adapter

}
