package webserver

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	es "core-interview/server/encryption_server"
	"fmt"
)

func Start() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/store", store).Methods("POST")
	rtr.HandleFunc("/retrieve/{id}/{key}", retrieve).Methods("GET")

	http.Handle("/", rtr)
	
	log.Print("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func store(w http.ResponseWriter, r *http.Request) {
	log.Print("Received POST on /store")
	vars := mux.Vars(r)
	id, data := vars["id"], vars["data"]
	log.Printf("--- from id=`%s`", id)
	aesKey, err := es.Store([]byte(id), []byte(data))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", string(aesKey))
}

func retrieve(w http.ResponseWriter, r *http.Request) {
	log.Print("Received GET on /retrieve")
	vars := mux.Vars(r)
	id, key := vars["id"], vars["key"]
	log.Printf("--- with id=`%s` and key=`%s`", id, key)
	payload, err := es.Retrieve([]byte(id), []byte(key))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", string(payload))
}