package chess

import (
	"fmt"
	"strings"
)

const positionStartFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type Position struct {
	board *Board
}

func NewPosition(board *Board) *Position {
	return &Position{
		board: board,
	}
}

func NewPositionStart() (*Position, error) {
	position, err := NewPositionFromFEN(positionStartFEN)
	if err != nil {
		return nil, fmt.Errorf("NewPositionFromFEN(%q): %w", positionStartFEN, err)
	}

	return position, nil
}

// Parses position's FEN to the Position structure.
//
// FEN argument example: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1".
func NewPositionFromFEN(fen string) (*Position, error) {
	// Required FEN parts: board, active color, castling rights, en passant, half move clock and full move number.
	const fenPartsLenRequired = 6

	fenParts := strings.Split(strings.TrimSpace(fen), " ")
	if len(fenParts) != fenPartsLenRequired {
		return nil, fmt.Errorf("FEN parts len required %d but got %d", fenPartsLenRequired, len(fenParts))
	}

	board, err := NewBoardFromFEN(fenParts[0])
	if err != nil {
		return nil, fmt.Errorf("NewBoardFromFen(%q): %w", fenParts[0], err)
	}

	return NewPosition(board), nil
}
