package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const apiURL = "https://api.openai.com/v1/chat/completions"

type GPTRequest struct {
	Messages []GPTMessage `json:"messages"`
	Model    string       `json:"model"`
}

type GPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPTResponse struct {
	Choices []struct {
		Message GPTMessage `json:"message"`
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
	requestBody, err := json.Marshal(GPTRequest{
		Messages: []GPTMessage{
			{Role: "user", Content: fmt.Sprintf("Write the value %s in full.", value)},
		},
		Model: "gpt-3.5-turbo",
	})
	if err != nil {
		fmt.Printf("Error creating request body: %v\n", err)
		return
	}

	//Create request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	var gptResponse GPTResponse
	err = json.Unmarshal(body, &gptResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return
	}

	if len(gptResponse.Choices) > 0 {
		fmt.Printf("Answer: %s\n", gptResponse.Choices[0].Message.Content)
	} else {
		fmt.Println("No answer found.")
	}
}