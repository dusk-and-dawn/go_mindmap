// main.go
package main

import (
	"encoding/json"
	"fmt"
	ihttp "go_mindmap/adapters/inbound/http"
	"go_mindmap/core/usecases"
	"log"
	"net/http"
)

type Node struct {
	Id       string
	Content  string
	Children []string
	X        float64
	Y        float64
}

func (f Node) sayHi() {
	fmt.Printf("Hi %v", f)
}

func getHandler() {

}

func main() {

	fmt.Println("In main angekommen")

	mux := http.NewServeMux()

	mh := ihttp.MapHandler{Mastruct: usecases.MapActions{Prefix: "dependency injected "}}

	mux.Handle("/map/{id}", mh)

	testNode := Node{
		Id:       "01",
		Content:  "Test-Node",
		Children: []string{"child1", "child2"},
		X:        260,
		Y:        260,
	}

	//testNode.sayHi()

	//encoder := json.NewEncoder(os.Stdout)

	barray, err := json.Marshal(testNode)

	if err != nil {
		fmt.Println(err)
		return
	}
	println(string(barray))

	//setupHandler(testNode)

	//http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
