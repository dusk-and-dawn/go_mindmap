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
	Id      string `json:"id"`
	Content string `json:"content"`
}

type HTTPNodeList struct {
	Nodes []HTTPNode `json:"nodes"`
}

func (f MapHandler) GetNode(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	mmap, mmap_bool := f.MAusecase.GetNodeByID(id) // Roman fragen wie ich dies anpassen muss um JSON zurück zu geben

	if !mmap_bool {
		http.Error(w, "Keine Node gefunden mit der ID", http.StatusNotFound)
		return
	}
	httpMM := HTTPNode{
		Id:      mmap.Id,
		Content: mmap.Content,
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
	fmt.Println(newNode.Content)
	node2neo := domain.Node{Id: newNode.Id, Content: newNode.Content}
	f.MAusecase.SaveNode(node2neo)

}

func (f MapHandler) GetNodes(w http.ResponseWriter, r *http.Request) {
	err := f.MAusecase.GetNodes()

	if err != nil {
		fmt.Println("kaputt aus GetNodes Http")
		return
	}
	snarkyComment, err := json.Marshal("check the Terminal")

	if err != nil {
		fmt.Println(err)
	}
	w.Write(snarkyComment)
	return

}

func (f MapHandler) ShowWaiting(w http.ResponseWriter, r *http.Request) {
	msg := "Hier noch nix im Frontend"

	barrayMsg := []byte(msg)
	w.Write(barrayMsg)
}

func (f MapHandler) ConnectNodes(w http.ResponseWriter, r *http.Request) {
	var newRelationship string
	// id1 := r.PathValue("id1")
	// id2 := r.PathValue("id2")
	err := json.NewDecoder(r.Body).Decode(&newRelationship)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "JSON nicht entschlüsselt bekommen", http.StatusBadRequest)
		return
	}
	// f.MAusecase.ConnectNodesFE(id1, id2)
}
