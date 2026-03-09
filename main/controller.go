package main

import (
	"encoding/json"
	"net/http"
)

func setupHandler(testNode Node) {
	http.HandleFunc("/mindmap", func(w http.ResponseWriter, _ *http.Request) {
		barray, err := json.Marshal(testNode)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Write(barray)
	})
}
