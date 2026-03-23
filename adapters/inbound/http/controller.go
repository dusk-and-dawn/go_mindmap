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

	mmap, mmap_bool := f.Mastruct.GetMapByID(id) // Roman fragen wie ich dies anpassen muss um JSON zurück zu geben
	if !mmap_bool {
		http.Error(w, "Keine Node gefunden mit der ID", http.StatusNotFound)
		return
	} else {
		barrayID := []byte(mmap.Content)
		w.Write(barrayID)
	}

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
