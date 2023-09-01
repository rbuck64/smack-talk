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

	baseURLHandler := func(w http.ResponseWriter, r *http.Request) {
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
	}

	getSmackTalkHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Print("HTMX request received!")
		log.Print(r.Header.Get("HX-Request"))

		htmlStr := fmt.Sprintf("<h2>This is a random fact from OpenAI</h2>")
		tmpl, _ := template.New("t").Parse(htmlStr)

		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", baseURLHandler)
	http.HandleFunc("/get-smack-talk/", getSmackTalkHandler)

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
