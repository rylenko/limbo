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
	enPassantSquare *Square
	halfMoveClock   uint8
	fullMoveNumber  uint16
}

// NewPosition creates a new position with passed parameters.
func NewPosition(
	board *Board,
	activeColor Color,
	castlingRights CastlingRights,
	enPassantSquare *Square,
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
	fenParts := strings.Split(fen, " ")
	if len(fenParts) != positionFENPartsCount {
		return nil, fmt.Errorf("FEN parts required %d but got %d", positionFENPartsCount, len(fenParts))
	}

	board, err := NewBoardFromFEN(fenParts[0])
	if err != nil {
		return nil, fmt.Errorf("NewBoardFromFEN(%q): %w", fenParts[0], err)
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

	halfMoveClockUint64, err := strconv.ParseUint(fenParts[4], 10, 8)
	if err != nil {
		return nil, fmt.Errorf("half move clock is not uint8: %w", err)
	}
	halfMoveClock := uint8(halfMoveClockUint64)

	fullMoveNumberUint64, err := strconv.ParseUint(fenParts[5], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("full move number is not uint16: %w", err)
	}
	fullMoveNumber := uint16(fullMoveNumberUint64)

	return NewPosition(board, activeColor, castlingRights, enPassantSquare, halfMoveClock, fullMoveNumber), nil
}

// CalculateMoves calculates all possible moves in the current position.
//
// TODO: test.
func (position *Position) CalculateMoves() []Move {
	// TODO: generate default moves and castlings.

	var moves []Move

	for _, piece := range NewPiecesOfColor(position.activeColor) {
		originBitboard, ok := position.board.bitboards[piece]
		if !ok || originBitboard == 0 {
			continue
		}

		for _, origin := range originBitboard.GetSquares() {
			pieceMoves := position.CalculatePieceMoves(piece, origin)
			moves = append(moves, pieceMoves...)
		}
	}

	return moves
}

// CalculatePieceMoves calculates all possible piece moves in the current position from passed origin.
//
// TODO: test.
func (position *Position) CalculatePieceMoves(piece Piece, origin Square) []Move {
	// TODO: generate default moves and castlings.

	if piece.Color() != position.activeColor {
		return nil
	}

	colorBitboard := position.board.GetColorBitboard(position.activeColor)
	pieceMoveBitboard := piece.Role().GetMoveBitboard(origin, piece.Color(), position.enPassantSquare)

	destBitboard := ^colorBitboard | pieceMoveBitboard
	destSquares := destBitboard.GetSquares()

	moves := make([]Move, 0, len(destSquares))

	for _, dest := range destSquares {
		isPromo := (piece == PieceWhitePawn && dest.Rank() == Rank8) ||
			(piece == PieceBlackPawn && dest.Rank() == Rank1)

		moves = append(moves, NewMove(origin, dest, isPromo))
	}

	return moves
}
