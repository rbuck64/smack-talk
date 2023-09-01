package main

type Member struct {
	DisplayName     string `json:"displayName"`
	ID              string `json:"id"`
	IsLeagueManager bool   `json:"isLeagueManager"`
}

type Team struct {
	Abbrev   string   `json:"abbrev"`
	ID       int      `json:"id"`
	Location string   `json:"location"`
	Nickname string   `json:"nickname"`
	Owners   []string `json:"owners"`
}

type LeagueStatus struct {
	CurrentMatchupPeriod int  `json:"currentMatchupPeriod"`
	IsActive             bool `json:"isActive"`
	LatestScoringPeriod  int  `json:"latestScoringPeriod"`
}

type LeagueSettings struct {
	Name string `json:"name"`
}

type LeagueAPIResponse struct {
	GameID          int            `json:"gameId"`
	ID              int            `json:"id"`
	Members         []Member       `json:"members"`
	ScoringPeriodID int            `json:"scoringPeriodId"`
	SeasonID        int            `json:"seasonId"`
	SegmentID       int            `json:"segmentId"`
	Settings        LeagueSettings `json:"settings"`
	Status          LeagueStatus   `json:"status"`
	Teams           []Team         `json:"teams"`
}

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
