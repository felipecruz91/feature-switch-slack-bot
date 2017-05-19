package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheckEndpoint returns 200 OK
func HealthCheckEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", HealthCheckEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
