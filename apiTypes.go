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
