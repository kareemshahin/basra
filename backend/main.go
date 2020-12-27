package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	router := mux.NewRouter()
	router.HandleFunc("/games", getGames).Methods("GET")
	router.HandleFunc("/games", createGame).Methods("POST")

	http.ListenAndServe(":8000", router)
}

func getGames(w http.ResponseWriter, r *http.Request) {
	games = GetGames()

	fmt.Println(games)
	respondWithJSON(w, http.StatusOK, games)
}

func createGame(w http.ResponseWriter, r *http.Request) {
	var gameReq Game
	err := json.NewDecoder(r.Body).Decode(&gameReq)

	if err != nil {
		respondWithError(w, 400, err.Error())
	}
	newGame, createErr := CreateGame(gameReq.Name)

	if createErr != nil {
		respondWithError(w, 400, createErr.Error())
	}
	respondWithJSON(w, http.StatusOK, newGame)
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
