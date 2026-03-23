package http

import (
	"encoding/json"
	"go_mindmap/core/usecases"
	"net/http"
)

type MapHandler struct {
	Mastruct usecases.MapActions
}

type HTTPMindmap struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

func (f MapHandler) GetMap(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	mmap, mmap_bool := f.Mastruct.GetMapByID(id) // Roman fragen wie ich dies anpassen muss um JSON zurück zu geben

	if !mmap_bool {
		http.Error(w, "Keine Node gefunden mit der ID", http.StatusNotFound)
		return
	}
	httpMM := HTTPMindmap{
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

func (f MapHandler) CreateMap(w http.ResponseWriter, r *http.Request) {

}

// func setupHandler(testNode Node) {
// 	http.HandleFunc("/testCall", func(w http.ResponseWriter, _ *http.Request) {

// 		barray, err := json.Marshal(testNode)
// 		if err != nil {
// 			w.WriteHeader(500)
// 			return
// 		}
// 		w.Write(barray)
// 	})
// }
