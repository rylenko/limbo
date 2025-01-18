package position

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/brunoga/deep"
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

// Copy deeply copies current position.
func (position *Position) DeepCopy() (*Position, error) {
	return deep.Copy(position)
}

// MoveRaw makes a raw move in the current position.
//
// Note that the move is raw, so it can, for example, put the active color in check.
//
// TODO: test.
func (position *Position) MoveRaw(move Move) error {
	if err := position.board.MoveRaw(move); err != nil {
		return fmt.Errorf("board.MoveRaw(%+v): %w", move, err)
	}

	if err := position.updateActiveColor(); err != nil {
		return fmt.Errorf("updateActiveColor(): %w", err)
	}

	if err := position.updateCastlingRightsRaw(move); err != nil {
		return fmt.Errorf("updateCastlingRightsRaw(%+v): %w", move, err)
	}

	if err := position.updateEnPassantSquareRaw(move); err != nil {
		return fmt.Errorf("updateEnPassantSquare(%+v): %w", move, err)
	}

	if err := position.updateHalfMoveClockRaw(move); err != nil {
		return fmt.Errorf("updateHalfMoveClockRaw(%+v): %w", move, err)
	}

	position.updateFullMoveNumber()

	return nil
}

// updateActiveColor updates active color to next active color.
//
// TODO: test.
func (position *Position) updateActiveColor() error {
	newColor, err := position.activeColor.Opposite()
	if err != nil {
		return fmt.Errorf("%s.Opposite(): %w", position.activeColor, err)
	}

	position.activeColor = newColor

	return nil
}

// updateCastlingRightsRaw updates castling rights depending on the passed move.
//
// Note that the move is raw, so it can, for example, put the active color in check.
//
// TODO: test.
func (position *Position) updateCastlingRightsRaw(move Move) error {
	originPiece, err := position.board.GetPieceFromSquare(move.origin)
	if err != nil {
		return fmt.Errorf("GetPieceFromSquare(%s): %w", move.origin, err)
	}

	if originPiece == PieceNil {
		return fmt.Errorf("no piece on the origin %s", move.origin)
	}

	var colorSideToDelete ColorSide

	if originPiece == PieceWhiteKing || move.origin == SquareA1 || move.dest == SquareA1 {
		colorSideToDelete = ColorSideWhiteQueen
	}

	if originPiece == PieceWhiteKing || move.origin == SquareH1 || move.dest == SquareH1 {
		colorSideToDelete = ColorSideWhiteKing
	}

	if originPiece == PieceBlackKing || move.origin == SquareA8 || move.dest == SquareA8 {
		colorSideToDelete = ColorSideBlackQueen
	}

	if originPiece == PieceBlackKing || move.origin == SquareH8 || move.dest == SquareH8 {
		colorSideToDelete = ColorSideBlackKing
	}

	position.castlingRights = slices.DeleteFunc(position.castlingRights, func(colorSide ColorSide) bool {
		return colorSide == colorSideToDelete
	})

	return nil
}

// updateEnPassantSquareRaw updates En Passant square depending on the passed move.
//
// Note that the move is raw, so it can, for example, put the active color in check.
//
// TODO: test.
func (position *Position) updateEnPassantSquareRaw(move Move) error {
	piece, err := position.board.GetPieceFromSquare(move.origin)
	if err != nil {
		return fmt.Errorf("GetPieceFromSquare(%s): %w", move.origin, err)
	}

	if piece == PieceNil {
		return fmt.Errorf("no piece on the origin %s", move.origin)
	}

	role, err := piece.Role()
	if err != nil {
		return fmt.Errorf("%s.Role(): %w", piece, err)
	}

	if role != RolePawn {
		return nil
	}

	originRank, err := move.origin.Rank()
	if err != nil {
		return fmt.Errorf("%s.Rank(): %w", move.origin, err)
	}

	destRank, err := move.dest.Rank()
	if err != nil {
		return fmt.Errorf("%s.Rank(): %w", move.dest, err)
	}

	enPassantSquareFile, err := move.origin.File()
	if err != nil {
		return fmt.Errorf("%s.File(): %w", move.origin, err)
	}

	var enPassantSquareRank Rank

	switch {
	case position.activeColor == ColorWhite && originRank == Rank2 && destRank == Rank4:
		enPassantSquareRank = Rank3
	case position.activeColor == ColorBlack && originRank == Rank7 && destRank == Rank5:
		enPassantSquareRank = Rank6
	default:
		return nil
	}

	square, err := NewSquare(enPassantSquareRank, enPassantSquareFile)
	if err != nil {
		return fmt.Errorf("NewSquare(%s, %s): %w", enPassantSquareRank, enPassantSquareFile, err)
	}

	position.enPassantSquare = square

	return nil
}

// updateFullMoveNumber updates full move number.
//
// TODO: test.
func (position *Position) updateFullMoveNumber() {
	if position.activeColor == ColorBlack {
		position.fullMoveNumber++
	}
}

// updateHalfMoveClock updates half move clock.
//
// Note that the move is raw, so it can, for example, put the active color in check.
//
// TODO: test.
func (position *Position) updateHalfMoveClockRaw(move Move) error {
	originPiece, err := position.board.GetPieceFromSquare(move.origin)
	if err != nil {
		return fmt.Errorf("GetPieceFromSquare(%s): %w", move.origin, err)
	}

	if originPiece == PieceNil {
		return fmt.Errorf("no piece on the origin %s", move.origin)
	}

	originPieceRole, err := originPiece.Role()
	if err != nil {
		return fmt.Errorf("%s.Role(): %w", originPiece, err)
	}

	if originPieceRole == RolePawn || move.tags.Contains(MoveTagCapture) {
		return nil
	}

	position.halfMoveClock++

	return nil
}
