// file mux_ex.go
package main

import (
	"github.com/gorilla/mux"

	"io/ioutil"
	"log"
	"net/http"

	"codetest/lottery"
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
	fileContents, _ := ioutil.ReadFile("./json/response.json")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(fileContents)
}

func main() {
	r := mux.NewRouter()
	log.Println("Waiting...")
	r.HandleFunc("/", baseHandler)
	r.Handle("/lottery", lottery.Get()).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))
}
