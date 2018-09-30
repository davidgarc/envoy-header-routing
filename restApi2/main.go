package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHelloWorld(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	res.Write([]byte("Hello from REST ** 2 **"))
	return
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/helloworld", GetHelloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":8002", router))
}
