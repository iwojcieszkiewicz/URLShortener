package main

import (
	"fmt"
	"net/http"
	"url-shortener/handler"
	"url-shortener/store"

	"github.com/gorilla/mux"
)

func main() {
	store := store.New()
	handler := handler.New(store)
	router := mux.NewRouter()

	router.HandleFunc("/shorten", handler.Shorten).Methods("POST")
	router.HandleFunc("/{code}", handler.Redirect).Methods("GET")

	fmt.Println("Serwer działa na http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
