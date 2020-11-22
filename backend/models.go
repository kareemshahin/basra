package main

type Game struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

type Player struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string
	Score int `json:"score"`
}
