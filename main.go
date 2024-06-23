package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/number_facts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := getNumberFact(r)
		err := json.NewEncoder(w).Encode(response)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}

func getNumberFact(r *http.Request) CustomResponse {
	query := r.URL.Query().Get("query")
	parsed, err := strconv.Atoi(query)

	if err != nil {
		resp := CustomResponse{Message: "Invalid query"}
		return resp
	}

	fact := getFactFromNumbersApi(parsed)
	resp := CustomResponse{Message: fact}
	return resp
}

func getFactFromNumbersApi(number int) string {
	num64 := int64(number)
	resp, err := http.Get("http://numbersapi.com/" + strconv.FormatInt(num64, 10) + "/trivia" + "?json")

	if err != nil {
		fmt.Println(err)
		return "Error getting fact"
	}

	defer resp.Body.Close()

	body := resp.Body
	fact, err := io.ReadAll(body)

	if err != nil {
		fmt.Println(err)
		return "Error reading fact"
	}

	numberFact := NumberFactsResponse{}
	err = json.Unmarshal(fact, &numberFact)

	if err != nil {
		return "Error reading fact"
	}

	return numberFact.Text
}

type NumberFactsResponse struct {
	Text string `json:"text"`
}

type CustomResponse struct {
	Message string `json:"message"`
}
