package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"message": "It's alive!"}`))
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"message": "get called"}`))
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte(`{"message": "post called"}`))
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, err := w.Write([]byte(`{"message": "put called"}`))
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"message": "delete called"}`))
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte(`{"message": "not found"}`))
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", health).Methods(http.MethodGet)
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
