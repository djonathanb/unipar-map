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

func routeHandler(w http.ResponseWriter, r *http.Request) {

}

func routeUtilitiesHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/environments", environmentsHandler)
	http.HandleFunc("/route", routeHandler)
	http.HandleFunc("/route/utilities", routeUtilitiesHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
