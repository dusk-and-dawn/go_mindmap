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

// type Node struct {
// 	Id       string
// 	Content  string
// 	Children []string
// 	X        float64
// 	Y        float64
// }

// wie würde ich diese function wieder nutzbar machen wenn ich die struct definition in map-actions machen will?
// func (f usecases.Node) sayHi() {
// 	fmt.Printf("Hi %v", f)
//  }

func getHandler() {

}

func main() {

	fmt.Println("In main angekommen")

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

	// a = usecases.Node{Id: "01", Content: "Super coole Testnode"}
	nodes := map[string]usecases.Node{}

	for _, value := range c {
		nodes[value.Id] = value
	}

	mh := ihttp.MapHandler{Mastruct: usecases.MapActions{Nodes: nodes}}

	mux.HandleFunc("GET /map/{id}", mh.GetMap)

	mux.HandleFunc("POST /map", mh.CreateMap)

	// testNode := Node{
	// 	Id:       "01",
	// 	Content:  "Test-Node",
	// 	Children: []string{"child1", "child2"},
	// 	X:        260,
	// 	Y:        260,
	// }

	//testNode.sayHi()

	//encoder := json.NewEncoder(os.Stdout)

	// barray, err := json.Marshal(testNode)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// println(string(barray))

	//setupHandler(testNode)

	//http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
