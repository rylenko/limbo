package chess

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// Required FEN parts: board, active color, castling rights, en passant, half move clock and full move number.
	positionFENPartsCount = 6
	positionStartFEN      = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

// Position represents the state of the game at a certain point in time.
type Position struct {
	board           *Board
	activeColor     Color
	castlingRights  CastlingRights
	enPassantSquare Square
	halfMoveClock   uint8
	fullMoveNumber  uint16
}

// NewPosition creates a new position with passed parameters.
func NewPosition(
	board *Board,
	activeColor Color,
	castlingRights CastlingRights,
	enPassantSquare Square,
	halfMoveClock uint8,
	fullMoveNumber uint16,
) *Position {
	return &Position{
		board:           board,
		activeColor:     activeColor,
		castlingRights:  castlingRights,
		enPassantSquare: enPassantSquare,
		halfMoveClock:   halfMoveClock,
		fullMoveNumber:  fullMoveNumber,
	}
}

// NewPositionStart creates game start position.
func NewPositionStart() (*Position, error) {
	position, err := NewPositionFromFEN(positionStartFEN)
	if err != nil {
		return nil, fmt.Errorf("NewPositionFromFEN(%q): %w", positionStartFEN, err)
	}

	return position, nil
}

// NewPositionFromFEN parses FEN to the Position structure.
//
// FEN argument example: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1".
func NewPositionFromFEN(fen string) (*Position, error) {
	parts := strings.Split(fen, " ")
	if len(parts) != positionFENPartsCount {
		return nil, fmt.Errorf("FEN parts required %d but got %d", positionFENPartsCount, len(parts))
	}

	board, err := NewBoardFromFEN(parts[0])
	if err != nil {
		return nil, fmt.Errorf("NewBoardFromFEN(%q): %w", parts[0], err)
	}

	activeColor, err := NewColorFromFEN(parts[1])
	if err != nil {
		return nil, fmt.Errorf("NewColorFromFEN(%q): %w", parts[1], err)
	}

	castlingRights, err := NewCastlingRightsFromFEN(parts[2])
	if err != nil {
		return nil, fmt.Errorf("NewCastlingRightsFromFEN(%q): %w", parts[2], err)
	}

	enPassantSquare, err := NewSquareEnPassantFromFEN(parts[3])
	if err != nil {
		return nil, fmt.Errorf("NewSquareEnPassantFromFEN(%q): %w", parts[3], err)
	}

	halfMoveClockUint64, err := strconv.ParseUint(parts[4], 10, 8)
	if err != nil {
		return nil, fmt.Errorf("ParseUint(%s, 10, 8): %w", parts[4], err)
	}
	halfMoveClock := uint8(halfMoveClockUint64)

	fullMoveNumberUint64, err := strconv.ParseUint(parts[5], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("ParseUint(%s, 10, 16): %w", parts[5], err)
	}
	fullMoveNumber := uint16(fullMoveNumberUint64)

	return NewPosition(board, activeColor, castlingRights, enPassantSquare, halfMoveClock, fullMoveNumber), nil
}

func (position *Position) Move(_ Move) (*Position, error) {
	newPosition := position
	return newPosition, nil
}
