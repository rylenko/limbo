package chess

import (
	"fmt"
	"strings"
)

// Position represents the state of the game at a certain point in time.
type Position struct {
	board           *Board
	activeColor     Color
	castlingRights  CastlingRights
	enPassantSquare *Square
}

// NewPosition creates a new position with passed parameters.
func NewPosition(board *Board, activeColor Color, castlingRights CastlingRights, enPassantSquare *Square) *Position {
	return &Position{
		board:           board,
		activeColor:     activeColor,
		castlingRights:  castlingRights,
		enPassantSquare: enPassantSquare,
	}
}

// NewPositionFromFEN parses FEN to the Position structure.
//
// FEN argument example: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1".
func NewPositionFromFEN(fen string) (*Position, error) {
	// Required FEN parts: board, active color, castling rights, en passant, half move clock and full move number.
	const fenPartsLenRequired = 6

	fenParts := strings.Split(fen, " ")
	if len(fenParts) != fenPartsLenRequired {
		return nil, fmt.Errorf("FEN parts len required %d but got %d", fenPartsLenRequired, len(fenParts))
	}

	board, err := NewBoardFromFEN(fenParts[0])
	if err != nil {
		return nil, fmt.Errorf("NewBoardFromFen(%q): %w", fenParts[0], err)
	}

	activeColor, err := NewColorFromFEN(fenParts[1])
	if err != nil {
		return nil, fmt.Errorf("NewColorFromFEN(%q): %w", fenParts[1], err)
	}

	castlingRights, err := NewCastlingRightsFromFEN(fenParts[2])
	if err != nil {
		return nil, fmt.Errorf("NewCastlingRightsFromFEN(%q): %w", fenParts[2], err)
	}

	enPassantSquare, err := NewSquareEnPassantFromFEN(fenParts[3])
	if err != nil {
		return nil, fmt.Errorf("NewSquareEnPassantFromFEN(%q): %w", fenParts[3], err)
	}

	return NewPosition(board, activeColor, castlingRights, enPassantSquare), nil
}
