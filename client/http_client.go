package client

import (
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"errors"
)

const SERVER = "http://localhost:8080"

type HttpClient struct{}

type Body struct {
	Id   string `json:"id"`
	Data string `json:"data"`
	Key  string `json:"key"`
}

func (h *HttpClient) Store(id, payload []byte) (aesKey []byte, err error) {
	m := Body{string(id), string(payload), ""}
	j, _ := json.Marshal(m)
	aesKey, err = post("store", j)
	return
}

func (h *HttpClient) Retrieve(id, aesKey []byte) (payload []byte, err error) {
	m := Body{string(id), "", string(aesKey)}
	j, _ := json.Marshal(m)
	payload, err = post("retrieve", j)
	return
}

func post(command string, jsonStr []byte) ([]byte, error) {
	url := SERVER + "/" + command

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
