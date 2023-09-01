package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

type OpenAICompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

type OpenAICompletionRequestBody struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Temperature float64 `json:"temperature"`
}

/*
curl https://api.openai.com/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer \
  -d '{
     "model": "gpt-3.5-turbo",
     "messages": [{"role": "user", "content": "Say this is a test!"}],
     "temperature": 0.7
   }'
*/

func main() {
	fmt.Println("Hello World!")

	// // https://fantasy.espn.com/apis/v3/games/ffl/leagueHistory/1868315?seasonId=2022&view=mTeam
	// getLeaugeMembersHandler := func(w http.ResponseWriter, r *http.Request) {
	// }
	// http.HandleFunc("/leagueMembers", getLeaugeMembersHandler)

	h1 := func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))

		res, error := http.Get("https://fantasy.espn.com/apis/v3/games/ffl/seasons/2023/segments/0/leagues/1868315")
		if error != nil {
			log.Fatal(error)
		}
		body, error := io.ReadAll(res.Body)

		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if error != nil {
			log.Fatal(error)
		}

		// parse the JSON Data
		var leagueAPIResponse LeagueAPIResponse
		error = json.Unmarshal(body, &leagueAPIResponse)
		if error != nil {
			log.Fatal(error)
		}

		teams := map[string][]Member{
			"Teams": {{DisplayName: "Kowski Krew", ID: "asdlkfj;as", IsLeagueManager: true}},
		}

		teams["Teams"] = leagueAPIResponse.Members

		temp.Execute(w, teams)

		str := getOpenAICall()
		fmt.Println(str)
		temp.Execute(w, str)
	}

	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getOpenAICall() map[string]string {
	fmt.Println("Doing the API Call")
	requestBody := []byte(`{
		"model": "gpt-3.5-turbo",
		"messages": [{"role": "user", "content": "Give me a random fact."}],
		"temperature": 0.7
	  }`)

	client := &http.Client{}

	req, err := http.NewRequest("POST", openAIAPIEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+OPEN_AI_KEY)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	// parse the JSON Data
	var completionResponse OpenAICompletionResponse
	error := json.Unmarshal(body, &completionResponse)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(string(completionResponse.Choices[0].Message.Content))

	completion := map[string]string{
		"Completion": "Hello There!",
	}

	completion["Completion"] = completionResponse.Choices[0].Message.Content

	return completion
}
