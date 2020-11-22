package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var games = []Game{
	{
		ID:    1,
		Name:  "game 1",
		State: "new",
	},
	{
		ID:    2,
		Name:  "game 2",
		State: "in_progress",
	},
}

func main() {
	fmt.Println("Hello, World!")

	router := mux.NewRouter()
	router.HandleFunc("/games", getGames).Methods("GET")
	router.HandleFunc("/games", createGame).Methods("POST")

	http.ListenAndServe(":8000", router)
}

func getGames(w http.ResponseWriter, r *http.Request) {
	fmt.Println(games)
	respondWithJSON(w, http.StatusOK, games)
}

func createGame(w http.ResponseWriter, r *http.Request) {
	new_game := Game{
		ID:    3,
		Name:  "game 3",
		State: "new_game",
	}

	games = append(games, new_game)
	respondWithJSON(w, http.StatusOK, new_game)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
