package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github/com/fcmdias/CSVAnalysis/services/backend/pkg/web"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 30,
}

type QueryBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func QueryOpenAI(prompt string, filter string) (string, error) {
	log.Println("Querying OpenAI with prompt:", prompt)

	requestBody := QueryBody{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are a data scientist assistant, skilled in explaining complex data in a simple and clear way. Your output will be wrapped around <div>{content}</div> html element tags and bootstrap for styling.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshalling request body: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	token := os.Getenv("AIToken")
	if token == "" {
		return "", errors.New("API token not found in environment variables")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to OpenAI: %w", err)
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK status from OpenAI: %d, response: %s", resp.StatusCode, string(response))
	}

	return string(response), nil
}

// Define a struct that matches the structure of your JSON data
type RequestData struct {
	Data []struct {
		Year  int `json:"year"`
		Total int `json:"total"`
	} `json:"data"`
}

func AskAIHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AskAIHandler called")

	if r.Method != "POST" {
		log.Println("Unsupported method:", r.Method)
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	filter := r.URL.Query().Get("filter")
	switch filter {
	case "all", "", " ":
		filter = "electric and hybrid"
	case "electric", "hybrid":
		// keep the same
	default:
		log.Printf("Invalid filter provided: %s", filter)
		http.Error(w, "Invalid filter option", http.StatusBadRequest)
		return
	}

	prompt := fmt.Sprintf("The data represents the number of %s cars registered in the state of Washington. Analyse and summarise this data: %s", filter, jsonData)
	response, err := QueryOpenAI(prompt, filter)
	if err != nil {
		log.Printf("Error querying OpenAI: %v", err)
		http.Error(w, "Error querying AI service", http.StatusInternalServerError)
		return
	}

	type Message struct {
		Content string `json:"content"`
	}
	type Choice struct {
		Message `json:"message"`
	}
	type Response struct {
		Choices []Choice `json:"choices"`
	}

	var dataResponse Response
	if err := json.Unmarshal([]byte(response), &dataResponse); err != nil {
		log.Printf("Error unmarshalling OpenAI response: %v", err)
		http.Error(w, "Error processing AI response", http.StatusInternalServerError)
		return
	}

	if len(dataResponse.Choices) == 0 {
		log.Println("No content received from OpenAI")
		http.Error(w, "No content received from AI service", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if _, err := w.Write([]byte(dataResponse.Choices[0].Content)); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func main() {

	http.Handle("/byyear", web.EnableCORSMiddleware(http.HandlerFunc(AskAIHandler)))
	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
