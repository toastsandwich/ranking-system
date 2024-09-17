package utils

type ScoreData struct {
	Username string  `json:"username"`
	Country  string  `json:"country"`
	State    string  `json:"state"`
	Score    float64 `json:"score"`
}
