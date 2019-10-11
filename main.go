package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/djonathanb/unipar-map/unipar"
)

func environmentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys := make(map[string]string, len(unipar.Environments))
	for key, value := range unipar.Environments {
		keys[key] = value.Name()
	}
	data, _ := json.Marshal(keys)
	w.Write([]byte(data))
}

func main() {
	http.HandleFunc("/environments", environmentsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
