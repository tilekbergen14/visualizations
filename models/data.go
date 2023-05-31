package models

type Data struct {
	Artist Artist `json:"artist"`
	Date Date `json:"date"`
	Location Location `json:"location"`
	Relation Relation `json:"relations"`
	ErrorCode int
	ErrorMessage string
}
