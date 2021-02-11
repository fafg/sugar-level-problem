package models

type User struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Samples []Sample `json:"samples"`
}

type Sample struct {
	Time  string  `json:"time"`
	Value float32 `json:"value"`
}

type SugarReport struct {
	Low    SugarResult  `json:"low"`
	Normal SugarResult  `json:"normal"`
	High   SugarResult  `json:"high"`
}

type SugarResult struct {
	Count   int    `json:"count"`
	Users []string `json:"users"` //array of user ids
}