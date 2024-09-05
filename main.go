package main

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
