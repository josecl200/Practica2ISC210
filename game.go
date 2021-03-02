package main

import (
	"encoding/json"
	"log"
)

const waitPaired = "Waiting to get paired"
const gameBegins = "Game begins!"
const draw = "Draw!"
const resetWaitPaired = "Opponent has been disconnected... Waiting to get paired again"
const winner = " wins! Congratulations!"

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type gameState struct {
	StatusMessage   string   `json:"statusMessage"`
	Fields          []field  `json:"fields"`
	PlayerSymbols   []string `json:"playerSymbols"`
	Started         bool     `json:"started"`
	Over            bool     `json:"over"`
	numberOfPlayers int
	PlayersTurn     int `json:"playersTurn"`
	numberOfMoves   int
}

type field struct {
	Set    bool   `json:"set"`
	Symbol string `json:"symbol"`
}

func newGameState() gameState {
	gs := gameState{
		StatusMessage:   waitPaired,
		Fields:          emptyFields(),
		PlayerSymbols:   []string{0: "X", 1: "O"},
		Started:         false,
		numberOfPlayers: 0,
		PlayersTurn:     0,
	}
	return gs
}

func emptyFields() []field {
	return []field{
		{}, {}, {},
		{}, {}, {},
		{}, {}, {},
	}
}

func (gs *gameState) addPlayer() {
	gs.numberOfPlayers++
	switch gs.numberOfPlayers {
	case 1:
		gs.StatusMessage = waitPaired
	case 2:
		gs.StatusMessage = gameBegins
		gs.Started = true
	}
}
func (gs *gameState) singlePlayerStart() {
	gs.numberOfPlayers = 1
	gs.StatusMessage = gameBegins
	gs.Started = true
}

func (gs *gameState) makeMove(playerNum int, moveNum int) {
	if moveNum <= 9 {
		if gs.isPlayersTurn(playerNum) {
			if gs.isLegalMove(moveNum) {
				gs.Fields[moveNum].Set = true
				gs.Fields[moveNum].Symbol = gs.PlayerSymbols[playerNum]
				gs.switchPlayersTurn(playerNum)
				gs.numberOfMoves++
				if won, symbol := gs.checkForWin(); won {
					gs.setWinner(symbol)
				} else {
					gs.checkForDraw()
				}
			}
		}
	} else {
		gs.specialMove(moveNum)
	}
}

func (gs *gameState) makeAIMoveHeuristic() {
	moves := []int{4, 0, 2, 6, 8, 1, 3, 5, 7}
	for _, val := range moves {
		if gs.isLegalMove(val) {
			gs.Fields[val].Set = true
			gs.Fields[val].Symbol = gs.PlayerSymbols[1]
			gs.switchPlayersTurn(1)
			gs.numberOfMoves++
			if won, symbol := gs.checkForWin(); won {
				gs.setWinner(symbol)
			} else {
				gs.checkForDraw()
			}
			break
		}
	}
}

func (gs *gameState) specialMove(moveNum int) {
	switch moveNum {
	case 10:
		gs.restartGame()
	}
}

func (gs *gameState) restartGame() {
	gs.StatusMessage = gameBegins
	gs.Fields = emptyFields()
	gs.Over = false
	gs.numberOfMoves = 0
}

func (gs *gameState) evaluateAI() int {
	for row := 0; row < 7; row += 3 {
		if gs.Fields[row].Symbol == gs.Fields[row+1].Symbol && gs.Fields[row+1].Symbol == gs.Fields[row+2].Symbol {
			if gs.Fields[row].Symbol == "O" {
				return 10
			}
			return -10
		}
	}
	for col := 0; col < 3; col++ {
		if gs.Fields[col].Symbol == gs.Fields[col+3].Symbol && gs.Fields[col+3].Symbol == gs.Fields[col+6].Symbol {
			if gs.Fields[col].Symbol == "O" {
				return 10
			}
			return -10
		}
	}

	if gs.Fields[0].Symbol == gs.Fields[4].Symbol && gs.Fields[4].Symbol == gs.Fields[8].Symbol {
		if gs.Fields[4].Symbol == "O" {
			return 10
		}
		return -10
	}

	if gs.Fields[2].Symbol == gs.Fields[4].Symbol && gs.Fields[4].Symbol == gs.Fields[6].Symbol {
		if gs.Fields[4].Symbol == "O" {
			return 10
		}
		return -10
	}
	return 0
}

func (gs *gameState) miniMax(depth int, max bool) int {
	score := gs.evaluateAI()
	if score == 10 {
		return score
	}
	if score == -10 {
		return score
	}
	if gs.numberOfMoves == 9 {
		return 0
	}
	if max {
		best := -1000
		for i := 0; i < 9; i++ {
			if gs.isLegalMove(i) {
				gs.Fields[i].Set = true
				gs.Fields[i].Symbol = "X"
				best = MaxInt(best, gs.miniMax(depth+1, !max))
				gs.Fields[i].Set = false
				gs.Fields[i].Symbol = ""
			}

		}
		return best
	} else {
		best := +1000
		for i := 0; i < 9; i++ {
			if gs.isLegalMove(i) {
				gs.Fields[i].Set = true
				gs.Fields[i].Symbol = "O"
				best = MaxInt(best, gs.miniMax(depth+1, !max))
				gs.Fields[i].Set = false
				gs.Fields[i].Symbol = ""
			}
		}
		return best
	}
	return -1
}

func (gs *gameState) makeAIMoveMinMax() {
	bestVal := -1000
	bestMove := -1

	for i := 0; i < 9; i++ {
		if gs.isLegalMove(i) {
			gs.Fields[i].Set = true
			gs.Fields[i].Symbol = "X"
			newVal := gs.miniMax(0, false)
			gs.Fields[i].Set = false
			gs.Fields[i].Symbol = ""
			if newVal > bestVal {
				bestMove = i
				bestVal = newVal
			}
		}
	}

	gs.Fields[bestMove].Set = true
	gs.Fields[bestMove].Symbol = gs.PlayerSymbols[1]
	gs.switchPlayersTurn(1)
	gs.numberOfMoves++
	if won, symbol := gs.checkForWin(); won {
		gs.setWinner(symbol)
	} else {
		gs.checkForDraw()
	}
}

func (gs *gameState) resetGame() {
	gs.restartGame()
	gs.Started = false
	gs.StatusMessage = resetWaitPaired
}

func (gs *gameState) checkForWin() (bool, string) {
	//rows
	if gs.Fields[0].Symbol == gs.Fields[1].Symbol && gs.Fields[1].Symbol == gs.Fields[2].Symbol && gs.Fields[2].Symbol != "" {
		return true, gs.Fields[0].Symbol
	}
	if gs.Fields[3].Symbol == gs.Fields[4].Symbol && gs.Fields[4].Symbol == gs.Fields[5].Symbol && gs.Fields[5].Symbol != "" {
		return true, gs.Fields[3].Symbol
	}
	if gs.Fields[6].Symbol == gs.Fields[7].Symbol && gs.Fields[7].Symbol == gs.Fields[8].Symbol && gs.Fields[8].Symbol != "" {
		return true, gs.Fields[7].Symbol
	}

	//columns
	if gs.Fields[0].Symbol == gs.Fields[3].Symbol && gs.Fields[3].Symbol == gs.Fields[6].Symbol && gs.Fields[6].Symbol != "" {
		return true, gs.Fields[0].Symbol
	}
	if gs.Fields[1].Symbol == gs.Fields[4].Symbol && gs.Fields[4].Symbol == gs.Fields[7].Symbol && gs.Fields[7].Symbol != "" {
		return true, gs.Fields[1].Symbol
	}
	if gs.Fields[2].Symbol == gs.Fields[5].Symbol && gs.Fields[5].Symbol == gs.Fields[8].Symbol && gs.Fields[8].Symbol != "" {
		return true, gs.Fields[2].Symbol
	}

	//diagonals
	if gs.Fields[0].Symbol == gs.Fields[4].Symbol && gs.Fields[4].Symbol == gs.Fields[8].Symbol && gs.Fields[8].Symbol != "" {
		return true, gs.Fields[0].Symbol
	}
	if gs.Fields[2].Symbol == gs.Fields[4].Symbol && gs.Fields[4].Symbol == gs.Fields[6].Symbol && gs.Fields[6].Symbol != "" {
		return true, gs.Fields[2].Symbol
	}
	return false, ""
}

func (gs *gameState) setWinner(symbol string) {
	gs.StatusMessage = symbol + winner
	gs.Over = true
}

func (gs *gameState) checkForDraw() {
	if gs.numberOfMoves == 9 {
		gs.StatusMessage = draw
		gs.Over = true
	}
}

func (gs *gameState) isLegalMove(field int) bool {
	return !gs.Fields[field].Set
}

func (gs *gameState) isPlayersTurn(playerNum int) bool {
	return playerNum == gs.PlayersTurn
}

func (gs *gameState) switchPlayersTurn(playerNum int) {
	switch playerNum {
	case 0:
		gs.PlayersTurn = 1
	case 1:
		gs.PlayersTurn = 0
	}
}

func (gs *gameState) gameStateToJSON() []byte {
	json, err := json.Marshal(gs)
	if err != nil {
		log.Fatal("Error in marshalling json:", err)
	}
	return json
}
