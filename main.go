package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintf(writer, "hello")

	})

	fmt.Println("Listenin at 8080")
	http.ListenAndServe(":8080", router)
}
