// main.go
package main

import (
	// "encoding/json"
	"fmt"
	ihttp "go_mindmap/adapters/inbound/http"
	"go_mindmap/core/usecases"
	"log"
	"net/http"
)

func getHandler() {

}

func main() {

	fmt.Println("In main angekommen")

	conn := setUpDB()

	mux := http.NewServeMux()

	a := usecases.Node{
		Id:       "01",
		Content:  "mega test node",
		Children: []string{"a", "b"},
		X:        160,
		Y:        200}

	b := usecases.Node{
		Id:       "02",
		Content:  "beste Teste",
		Children: []string{"a", "b"},
		X:        160,
		Y:        200}

	c := []usecases.Node{a, b}

	commitMmapToDB(conn, a)
	var testStringID string
	var testStringContent string
	err := conn.QueryRow("SELECT * FROM mmaps WHERE id = ?", 1).Scan(&testStringID, &testStringContent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(testStringID, testStringContent)

	nodes := map[string]usecases.Node{}

	for _, value := range c {
		nodes[value.Id] = value
	}

	mh := ihttp.MapHandler{Mastruct: usecases.MapActions{Nodes: nodes}}

	mux.HandleFunc("GET /map/{id}", mh.GetMap)

	mux.HandleFunc("POST /makemap", mh.CreateMap)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
