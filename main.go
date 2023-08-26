package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	// // https://fantasy.espn.com/apis/v3/games/ffl/leagueHistory/1868315?seasonId=2022&view=mTeam
	// getLeaugeMembersHandler := func(w http.ResponseWriter, r *http.Request) {
	// }
	// http.HandleFunc("/leagueMembers", getLeaugeMembersHandler)

	h1 := func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))

		res, error := http.Get("https://fantasy.espn.com/apis/v3/games/ffl/leagueHistory/1868315?seasonId=2022")
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

		teams["Teams"] = leagueAPIResponse[0].Members

		temp.Execute(w, teams)
	}

	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
