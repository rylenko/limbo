package chess

import "fmt"

// Game represents chess game with all position history.
type Game struct {
	positions []*Position
}

// NewGame creates a new game with passed parameters.
func NewGame(positions []*Position) *Game {
	return &Game{
		positions: positions,
	}
}

// NewGameStart creates a start of the game.
func NewGameStart() (*Game, error) {
	position, err := NewPositionStart()
	if err != nil {
		return nil, fmt.Errorf("NewPositionStart(): %w", err)
	}

	positions := []*Position{position}

	return NewGame(positions), nil
}
