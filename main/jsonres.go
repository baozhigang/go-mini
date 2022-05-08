package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/returnjson", returnjson)

	log.Fatal(http.ListenAndServe(":80", nil))
}

func returnjson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "status created"
	resp["data"] = "it is ok"
	jsonRes, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("error happen %s", err)
	}
	w.Write(jsonRes)
	return
}
