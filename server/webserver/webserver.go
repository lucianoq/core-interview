package webserver

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	es "core-interview/server/encryption_server"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

func Start() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/store", store).Methods("POST")
	rtr.HandleFunc("/retrieve", retrieve).Methods("POST")

	http.Handle("/", rtr)

	log.Print("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func store(w http.ResponseWriter, r *http.Request) {
	log.Print("Received POST on /store")

	type StoreRequest struct {
		Id   string `json:"id"`
		Data string `json:"data"`
	}

	body, err := ioutil.ReadAll(r.Body);
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req StoreRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("--- from id=`%s`", req.Id)

	aesKey, err := es.Store(req.Id, req.Data)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", aesKey)
}

func retrieve(w http.ResponseWriter, r *http.Request) {
	log.Print("Received GET on /retrieve")

	type RetrieveRequest struct {
		Id  string `json:"id"`
		Key string `json:"key"`
	}
	body, err := ioutil.ReadAll(r.Body);
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req RetrieveRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("--- with id=`%s`", req.Id)
	payload, err := es.Retrieve(req.Id, req.Key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", payload)
}