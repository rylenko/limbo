package chess

import (
	"fmt"
	"slices"
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

// Makes a raw move, returns the updated state.
//
// Note that the move is raw, so it can, for example, put the active color in check.
//
// TODO: test
func (position *Position) MoveRaw(move Move) (*Position, error) {
	newBoard, err := position.board.MoveRaw(move)
	if err != nil {
		return nil, fmt.Errorf("board.MoveRaw(%s): %w", move, err)
	}

	newActiveColor, err := position.activeColor.Opposite()
	if err != nil {
		return nil, fmt.Errorf("%s.Opposite(): %w", position.activeColor, err)
	}

	newCastlingRights, err := position.updateCastlingRightsRaw(move)
	if err != nil {
		return nil, fmt.Errorf("updateCastlingRightsRaw(%s): %w", move, err)
	}

	newHalfMoveClock, err := position.updateHalfMoveClockRaw(move)
	if err != nil {
		return nil, fmt.Errorf("updateHalfMoveClockRaw(%s): %w", move, err)
	}

	newPosition := NewPosition(
		newBoard, newActiveColor, newCastlingRights, newEnPassantSquare, newHalfMoveClock, position.updateFullMoveNumber())

	return newPosition, nil
}

// updateCastlingRights updates castling rights depending on the passed move. Returns updated castling rights.
//
// Note that the move is raw, so it can, for example, put the active color in check.
//
// TODO: test
func (position *Position) updateCastlingRightsRaw(move Move) (CastlingRights, error) {
	newCastlingRights = position.castlingRights.clone()

	originPiece, err := position.board.GetSquarePiece(move.origin)
	if err != nil {
		return nil, fmt.Errorf("GetSquarePiece(%s): %w", move.origin, err)
	}

	var colorSideToDelete ColorSide

	if originPiece == WhiteKing || move.origin == SquareA1 || move.dest == SquareA1 {
		colorSideToDelete = ColorSideWhiteQueen
	}
	if originPiece == WhiteKing || move.origin == SquareH1 || move.dest == SquareH1 {
		colorSideToDelete = ColorSideWhiteKing
	}
	if originPiece == BlackKing || move.origin == SquareA8 || move.dest == SquareA8 {
		colorSideToDelete = ColorSideBlackQueen
	}
	if originPiece == BlackKing || move.origin == SquareH8 || move.dest == SquareH8 {
		colorSideToDelete = ColorSideBlackKing
	}

	newCastlingRights = slices.DeleteFunc(newCastlingRights, func(colorSide ColorSide) {
		return colorSide == colorSideToDelete
	})

	return newCastlingRights, nil
}

// updateFullMoveNumber updates full move number.
//
// TODO: test
func (position *Position) updateFullMoveNumber() uint64 {
	number := position.fullMoveNumber
	if position.activeColor == ColorBlack {
		number++
	}

	return number
}

// updateHalfMoveClock updates half move clock.
//
// Note that the move is raw, so it can, for example, put the active color in check.
//
// TODO: test
func (position *Position) updateHalfMoveClockRaw(move Move) (uint8, error) {
	destPiece, err := position.board.GetSquarePiece(move.dest)
	if err != nil {
		return 0, fmt.Errorf("GetSquarePiece(%s): %w", move.dest, err)
	}

	destPieceRole, err := destPiece.Role()
	if err != nil {
		return 0, fmt.Errorf("%s.Role(): %w", destPiece, err)
	}

	if destPieceRole == RolePawn || move.HasTag(Capture) {
		return 0, nil
	}

	return position.halfMoveClock + 1
}
