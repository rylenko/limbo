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
		moves = append(moves, position.CalculatePieceMoves(piece)...)
	}

	return moves
}

// CalculatePieceMoves calculates all possible piece moves in the current position from passed origin.
//
// TODO: test.
func (position *Position) CalculatePieceMoves(piece Piece) []Move {
	// TODO: generate default moves and castlings.

	if piece.Color() != position.activeColor {
		return nil
	}

	var moves []Move

	for _, origin := range position.board.bitboards[piece].GetSquares() {
		for _, rawDest := range position.getPieceRawMovesBitboard(piece, origin).GetSquares() {
			isPromo := (piece == PieceWhitePawn && rawDest.Rank() == Rank8) ||
				(piece == PieceBlackPawn && rawDest.Rank() == Rank1)

			moves = append(moves, NewMove(origin, rawDest, isPromo))
		}
	}

	return moves
}

// getHorizontalVerticalRawMovesBitboard gets horizontal and vertical move Bitboard from passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put them in checkmate.
//
// TODO: test.
func (position *Position) getHorizontalVerticalRawMovesBitboard(color Color, origin Square) Bitboard {
	rankBitboard := Bitboard(0).SetSquares(NewSquaresOfRank(origin.Rank())...)
	fileBitboard := Bitboard(0).SetSquares(NewSquaresOfFile(origin.File())...)

	return position.getLinearRawMovesBitboard(color, origin, rankBitboard) |
		position.getLinearRawMovesBitboard(color, origin, fileBitboard)
}

// getKingRawMovesBitboard gets king move Bitboard from passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put them in checkmate.
//
// TODO: test.
func (position *Position) getKingRawMovesBitboard(color Color, origin Square) Bitboard {
	colorBitboard := position.board.GetColorBitboard(color)
	return roleKingMoveBitboards[origin] & ^colorBitboard
}

// getKnightRawMovesBitboard gets knight move Bitboard from passed origin.
//
// Note that the moves are raw, that is, for example, the knight moves can put their king in checkmate.
//
// TODO: test.
func (position *Position) getKnightRawMovesBitboard(color Color, origin Square) Bitboard {
	colorBitboard := position.board.GetColorBitboard(color)
	return roleKnightMoveBitboards[origin] & ^colorBitboard
}

// getLinearRawMovesBitboard gets linear move Bitboard from passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put them in checkmate.
//
// TODO: test.
func (position *Position) getLinearRawMovesBitboard(color Color, origin Square, lineBitboard Bitboard) Bitboard {
	originBitboard := Bitboard(0).SetSquares(origin)
	occupiedLineBitboard := position.board.GetOccupiedBitboard() & lineBitboard

	movesToBlockerBitboard := lineBitboard & ((occupiedLineBitboard - 2*originBitboard) ^
		(occupiedLineBitboard.Reverse() - 2*originBitboard.Reverse()).Reverse())

	colorBitboard := position.board.GetColorBitboard(color)

	return movesToBlockerBitboard & ^colorBitboard
}

// getPawnRawMovesBitboard gets pawn moves Bitboard from passed origin.
//
// Note that the moves are raw, that is, for example, the pawn moves can put their king in checkmate.
//
// TODO: test.
func (position *Position) getPawnRawMovesBitboard(color Color, origin Square) Bitboard {
	if (color == ColorBlack && origin.Rank() == Rank1) || (color == ColorWhite && origin.Rank() == Rank8) {
		return 0
	}

	originBitboard := Bitboard(0).SetSquares(origin)
	unoccupiedBitboard := ^position.board.GetOccupiedBitboard()

	allCapturesBitboard := position.board.GetColorBitboard(color.Opposite())
	if position.enPassantSquare != nil {
		allCapturesBitboard = allCapturesBitboard.SetSquares(*position.enPassantSquare)
	}

	var bitboard Bitboard

	switch color {
	case ColorBlack:
		bitboard |= originBitboard << len(files) & unoccupiedBitboard

		if origin.Rank() == Rank7 {
			bitboard |= originBitboard << (2 * len(files)) & unoccupiedBitboard //nolint:mnd // Skip all files twice.
		}
		if origin.File() != FileA {
			bitboard |= originBitboard << (len(files) + 1) & allCapturesBitboard
		}
		if origin.File() != FileH {
			bitboard |= originBitboard << (len(files) - 1) & allCapturesBitboard
		}
	case ColorWhite:
		bitboard |= originBitboard >> len(files) & unoccupiedBitboard

		if origin.Rank() == Rank2 {
			bitboard |= originBitboard >> (2 * len(files)) & unoccupiedBitboard //nolint:mnd // Skip all files twice.
		}
		if origin.File() != FileA {
			bitboard |= originBitboard >> (len(files) - 1) & allCapturesBitboard
		}
		if origin.File() != FileH {
			bitboard |= originBitboard >> (len(files) + 1) & allCapturesBitboard
		}
	}

	return bitboard
}

// getPieceRawMovesBitboard gets piece move Bitboard from passed origin.
//
// Note that the moves are raw, that is, for example, the piece moves can put their king in checkmate.
//
// TODO: test.
func (position *Position) getPieceRawMovesBitboard(piece Piece, origin Square) Bitboard {
	if piece.Color() != position.activeColor {
		return 0
	}

	switch piece.Role() {
	case RoleKing:
		return position.getKingRawMovesBitboard(piece.Color(), origin)
	case RoleQueen:
		return 0
	case RoleRook:
		return position.getHorizontalVerticalRawMovesBitboard(piece.Color(), origin)
	case RoleBishop:
		return 0
	case RoleKnight:
		return position.getKnightRawMovesBitboard(piece.Color(), origin)
	case RolePawn:
		return position.getPawnRawMovesBitboard(piece.Color(), origin)
	default:
		return 0
	}
}
