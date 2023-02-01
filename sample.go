package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type foods struct {
	Name  string
	Price int
}

func Sample() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello worlds"))
	})

	r.HandleFunc("/test", test).Methods("GET")
	r.HandleFunc("/product/{id}", paramshandler).Methods("GET")
	r.HandleFunc("/users", quryhandler).Methods("GET")
	r.HandleFunc("/food", bodyhandler).Methods("POST")

	log.Println("app runn on port 3001")
	err := http.ListenAndServe(":3001", r)
	if err != nil {
		log.Fatal(err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	w.Write([]byte("hello test"))
}

func paramshandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)["id"]

	fmt.Fprintf(w, "params : %v", vars)
}

func quryhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := r.URL.Query()

	username := vars["username"][0]

	fmt.Fprintf(w, "query : %v", username)
}

func bodyhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var fod foods
	json.NewDecoder(r.Body).Decode(&fod)

	fmt.Println(fod)
}

// FE -> /auth -> checkUser -> passwordFromDb -> createToken <-
