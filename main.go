package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)


type ReverseRequest struct {
	Text string `json:"text"`
}


type ReverseResponse struct {
	Reversed string `json:"reversed"`
}

func main() {
	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/reverse", handleReverse)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePing(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(responseWriter).Encode(map[string]string{"message": "pong"})
}

func handleReverse(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil || len(body) == 0 {
		http.Error(responseWriter, "Invalid request body", http.StatusBadRequest)
		return
	}

	var req ReverseRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(responseWriter, "Invalid JSON", http.StatusBadRequest)
		return
	}

	reversed := reverseString(req.Text)

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(ReverseResponse{Reversed: reversed})
}


func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}