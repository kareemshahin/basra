package main

/*
	Module to handle games that are stored and maintained in memory
*/

import (
	"fmt"
	"github.com/google/uuid"
)

var games = make(map[uuid.UUID]Game)

// states
const NEW_GAME = "new_game"
const START_OF_ROUND = "start_of_round"
const ROUND_IN_PROGRESS = "round_in_progress"
const END_OF_ROUND = "end_of_round"
const COMPLETED = "completed"

// TODO: add creator to game
func CreateGame(name string) (Game, error) {
	newId, err := uuid.NewUUID()

	if err != nil {
		fmt.Println(err)
		return Game{}, err
	}

	newGame := Game{
		ID:      newId,
		Name:    name,
		State:   NEW_GAME,
		Players: []GamePlayer{},
	}

	games[newId] = newGame

	return newGame, err
}

func GetGames() map[uuid.UUID]Game {
	return games
}
