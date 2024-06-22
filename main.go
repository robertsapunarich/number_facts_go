package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/number_facts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getNumberFact(r)))
	})
	http.ListenAndServe(":8080", nil)
}

func getNumberFact(r *http.Request) string {
	query := r.URL.Query().Get("query")
	return "your query is " + query
}
