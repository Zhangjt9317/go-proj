package main

type Poll struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Options  []Option `json:"options"`
}

type Option struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Votes int    `json:"votes"`
}
