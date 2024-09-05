package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const apiURl = "https://api.openai.com/v1/completions"

type GTPRequest struct {
	Messages []GTPMessage `json:"messages"`
	Model    string       `json:"model"`
}

type GTPMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GTPResponse struct {
	Choices []struct {
		Message GTPMessage `json:"message"`
	} `json:"choices"`
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY not set")
		return
	}

	var value string
	fmt.Print("Enter a monetary value: ")
	fmt.Scanln(&value)

	//Create payload
	requestBody, err := json.Marshal(GTPRequest{
		Messages: []GTPMessage{
			{Role: "user", Content: fmt.Sprintf("Write the value %s in full.", value)},
		},
		Model: "gpt-3.5-turbo",
	})
	if err != nil {
		fmt.Printf("Error creating request body: %v\n", err)
		return
	}

	//Create request
	req, err := http.NewRequest("POST", apiURl, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}


	//Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	//ChatGPT api key
	client := &http.Client{}
	resp, err := client.Do(req) //Send request
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}

	defer resp.Body.Close()
}