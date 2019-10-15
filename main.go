package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/djonathanb/unipar-map/unipar"
)

func environmentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys := make(map[string]string, len(unipar.Environments))
	for key, value := range unipar.Environments {
		keys[key] = value.Name
	}
	data, _ := json.Marshal(keys)
	w.Write([]byte(data))
}

// RouteResponse response for route calling
type RouteResponse struct {
	// Path     []unipar.Environment `json:"path"`
	Path     []unipar.EnvironmentStep `json:"path"`
	Distance int                      `json:"distance"`
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Query())
	reg := regexp.MustCompile(`^/from/(?P<from>\w*)/to/(?P<to>\w*)/?`)
	m := reg.FindStringSubmatch(r.URL.String())

	from := m[1]
	to := m[2]
	fromEnv := unipar.Environments[from]
	toEnv := unipar.Environments[to]

	path, distance := unipar.MinimumPath(fromEnv, toEnv)
	response := RouteResponse{Path: path, Distance: distance}

	data, _ := json.Marshal(response)
	w.Write([]byte(data))
}

func routeUtilitiesHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Query())
	reg := regexp.MustCompile(`^/bathroom/(?P<from>\w*)/?`)
	m := reg.FindStringSubmatch(r.URL.String())

	from := m[1]

	fromEnv := unipar.Environments[from]

	path, distance := unipar.MinimumPathToBathroom(fromEnv)
	response := RouteResponse{Path: path, Distance: distance}

	data, _ := json.Marshal(response)
	w.Write([]byte(data))
}

func main() {
	http.HandleFunc("/environments", environmentsHandler)
	http.HandleFunc("/from/", routeHandler)
	http.HandleFunc("/bathroom/", routeUtilitiesHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
