package http

import (
	"encoding/json"
	"fmt"
	"go_mindmap/core/domain"
	"go_mindmap/core/usecases"
	"net/http"
)

type MapHandler struct {
	MAusecase usecases.MapActions
}

type HTTPNode struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type HTTPNodeList struct {
	Nodes []HTTPNode `json:"nodes"`
}

type HTTPRelationshipList struct {
	Relationships []HTTPRelationship `json:"relationships"`
}

type HTTPRelationship struct {
	NameA string `json:"nameA"`
	NameB string `json:"nameB"`
	Kind  string `json:"kind"`
	Color string `json:"color"`
}

func (f MapHandler) GetNode(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	mmap, mmap_bool := f.MAusecase.GetNodeByID(id) // Roman fragen wie ich dies anpassen muss um JSON zurück zu geben

	if !mmap_bool {
		http.Error(w, "Keine Node gefunden mit der ID", http.StatusNotFound)
		return
	}
	httpMM := HTTPNode{
		Id:   mmap.Id,
		Name: mmap.Name,
	}
	barrayID, err := json.Marshal(httpMM)

	if err != nil {
		http.Error(w, "something went wrong", 500)
		return
	}
	w.Write(barrayID)
}

func (f MapHandler) CreateNode(w http.ResponseWriter, r *http.Request) {
	var newNode HTTPNode
	err := json.NewDecoder(r.Body).Decode(&newNode)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "JSON nicht entschlüsselt bekommen", http.StatusBadRequest)
		return
	}
	fmt.Println(newNode.Name)
	node2neo := domain.Node{Id: newNode.Id, Name: newNode.Name}
	f.MAusecase.SaveNode(node2neo)

}

func (f MapHandler) GetNodes(w http.ResponseWriter, r *http.Request) {
	nodes, err := f.MAusecase.GetNodes()

	if err != nil {
		fmt.Println("kaputt aus GetNodes Http")
		http.Error(w, "could not fetch nodes", http.StatusInternalServerError)
		return
	}

	var httpNodes []HTTPNode
	for _, n := range nodes {
		httpNodes = append(httpNodes, HTTPNode{Id: n.Id, Name: n.Name})
	}

	response := HTTPNodeList{Nodes: httpNodes}
	payload, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "99 Problems", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (f MapHandler) GetConnections(w http.ResponseWriter, r *http.Request) {
	connections, err := f.MAusecase.GetConnections()

	if err != nil {
		fmt.Println("kaputt aus GetNodes Http")
		http.Error(w, "could not fetch nodes", http.StatusInternalServerError)
		return
	}

	var httpCons []HTTPRelationship
	for _, n := range connections {
		httpCons = append(httpCons, HTTPRelationship{NameA: n.NameA, NameB: n.NameB, Kind: n.Kind, Color: n.Color})
	}

	response := HTTPRelationshipList{Relationships: httpCons}
	payload, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "99 Problems", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (f MapHandler) ShowWaiting(w http.ResponseWriter, r *http.Request) {
	msg := "Hier noch nix im Frontend"

	barrayMsg := []byte(msg)
	w.Write(barrayMsg)
}

func (f MapHandler) ConnectNodes(w http.ResponseWriter, r *http.Request) {
	var newRelationship HTTPRelationship
	// id1 := r.PathValue("id1")
	// id2 := r.PathValue("id2")
	err := json.NewDecoder(r.Body).Decode(&newRelationship)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "JSON nicht entschlüsselt bekommen", http.StatusBadRequest)
		return
	}
	err = f.MAusecase.ConnectNodesFE(newRelationship.NameA, newRelationship.NameB, newRelationship.Kind, newRelationship.Color)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Connection failed between nodes", http.StatusInternalServerError)
		return
	}
}
