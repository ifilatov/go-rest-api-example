package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ifilatov/hello-go/api"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"message": "It's alive!"}`))
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func getPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	person := api.GetPerson(params["id"])
	// the best way to check for an empty Person
	if person.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(person)
		return
	}
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func getPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(api.GetPeople())
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func createPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person api.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(api.CreatePerson(person))
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func modifyPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person api.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	matched, people := api.ModifyPerson(person)
	if !matched {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(people)
		return
	}
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func deletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	matched, people := api.DeletePerson(params["id"])
	if !matched {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(people)
		return
	}
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func init() {
	api.CreatePerson(api.Person{ID: "1", Firstname: "Fred", Lastname: "Flintstone"})
	api.CreatePerson(api.Person{ID: "2", Firstname: "Wilma", Lastname: "Flintstone"})
	api.CreatePerson(api.Person{ID: "3", Firstname: "Barney", Lastname: "Rubble"})
	api.CreatePerson(api.Person{ID: "4", Firstname: "Betty", Lastname: "Rubble"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", health).Methods(http.MethodGet)
	r.HandleFunc("/people", getPeopleEndpoint).Methods("GET")
	r.HandleFunc("/people/{id}", getPersonEndpoint).Methods("GET")
	r.HandleFunc("/people/{id}", createPersonEndpoint).Methods("POST")
	r.HandleFunc("/people/{id}", modifyPersonEndpoint).Methods("PUT")
	r.HandleFunc("/people/{id}", deletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
