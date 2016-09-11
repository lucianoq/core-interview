package webserver

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	es "core-interview/server/encryption_server"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strings"
	"core-interview/server/storage"
	"strconv"
)

type Request struct {
	Id   string `json:"id"`
	Data string `json:"data"`
	Key  string `json:"key"`
}

func Start(port int) {
	err := storage.Check()
	if err != nil {
		log.Print(err.Error())
		return
	}

	rtr := mux.NewRouter()

	rtr.HandleFunc("/store", store).Methods("POST")
	rtr.HandleFunc("/retrieve", retrieve).Methods("POST")

	http.Handle("/", rtr)

	log.Println("Server starting...")
	defer log.Println("Server stopping.")
	defer storage.DBClose()

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}

func store(w http.ResponseWriter, r *http.Request) {
	log.Print("Received POST on /store")
	req, err := getVar(r)
	if err != nil || req.Id == "" || req.Data == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("- Correct request with id `%s`\n", req.Id)

	aesKey, err := es.Store(req.Id, req.Data)
	//log.Print("AES key = " + aesKey)
	exit(w, aesKey, err)
}

func retrieve(w http.ResponseWriter, r *http.Request) {
	log.Print("Received GET on /retrieve")
	req, err := getVar(r)
	if err != nil || req.Id == "" || req.Key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("- Correct request with id `%s`\n", req.Id)

	payload, err := es.Retrieve(req.Id, req.Key)

	exit(w, payload, err)
}

func exit(w http.ResponseWriter, output string, err error) {
	if err != nil {
		log.Print(err.Error())
		switch {
		case strings.HasPrefix(err.Error(), "UNIQUE"):
			w.WriteHeader(http.StatusConflict)
		case strings.HasPrefix(err.Error(), "sql: no rows"):
			w.WriteHeader(http.StatusNotFound)
		case strings.HasPrefix(err.Error(), "illegal base64"):
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	fmt.Fprintf(w, "%s", output)
}

func getVar(r *http.Request) (Request, error) {
	body, err := ioutil.ReadAll(r.Body);
	if err != nil {
		return Request{}, err
	}
	var req Request
	err = json.Unmarshal(body, &req)
	if err != nil {
		return Request{}, err
	}
	return req, nil
}