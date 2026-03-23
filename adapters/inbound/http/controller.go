package http

import (
	"go_mindmap/core/usecases"
	"net/http"
)

type MapHandler struct {
	Mastruct usecases.MapActions
}

func (f MapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	mmap := f.Mastruct.GetMapByID(id)
	barrayID := []byte(mmap)
	w.Write(barrayID)
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
