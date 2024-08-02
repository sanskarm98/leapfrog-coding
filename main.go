package main

import (
	"fmt"
	"leapfrog-coding/handlers"
	"leapfrog-coding/operations"
	"net/http"
)

func main() {

	matrixOps := operations.NewMatrixOps()

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		handlers.EchoHandler(w, r, matrixOps)
	})

	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {
		handlers.InvertHandler(w, r, matrixOps)
	})

	http.HandleFunc("/flatten", func(w http.ResponseWriter, r *http.Request) {
		handlers.FlattenHandler(w, r, matrixOps)
	})

	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		handlers.SumHandler(w, r, matrixOps)
	})

	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		handlers.MultiplyHandler(w, r, matrixOps)
	})

	// Start the server
	fmt.Println("Server started at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
