package chess

import "errors"

// PieceType represents chess piece types.
type PieceType int8

const (
	PieceTypeWhiteKing PieceType = iota
	PieceTypeWhiteQueen
	PieceTypeWhiteRook
	PieceTypeWhiteBishop
	PieceTypeWhiteKnight
	PieceTypeWhitePawn
	PieceTypeBlackKing
	PieceTypeBlackQueen
	PieceTypeBlackRook
	PieceTypeBlackBishop
	PieceTypeBlackKnight
	PieceTypeBlackPawn
)

// Mapping of FEN string to corresponding PieceType.
var pieceTypeFromFENMap = map[string]PieceType{
	"k": PieceTypeBlackKing,
	"q": PieceTypeBlackQueen,
	"r": PieceTypeBlackRook,
	"b": PieceTypeBlackBishop,
	"n": PieceTypeBlackKnight,
	"p": PieceTypeBlackPawn,
	"K": PieceTypeWhiteKing,
	"Q": PieceTypeWhiteQueen,
	"R": PieceTypeWhiteRook,
	"B": PieceTypeWhiteBishop,
	"N": PieceTypeWhiteKnight,
	"P": PieceTypeWhitePawn,
}

// NewPieceTypeFromFEN parses FEN to corresponding PieceType or returns an error.
//
// FEN argument examples: "k", "q", "r", "b", "n", "p" for black pieces and and the same, but in upper case, for whites.
func NewPieceTypeFromFEN(fen string) (PieceType, error) {
	pieceType, ok := pieceTypeFromFENMap[fen]
	if !ok {
		return pieceType, errors.New("unknown FEN")
	}
	return pieceType, nil
}

// NewPieceTypesFromColor returns slice of piece types of passed Color.
func NewPieceTypesFromColor(color Color) []PieceType {
	switch color {
	case ColorBlack:
		return []PieceType{
			PieceTypeBlackKing,
			PieceTypeBlackQueen,
			PieceTypeBlackRook,
			PieceTypeBlackBishop,
			PieceTypeBlackKnight,
			PieceTypeBlackPawn}
	case ColorWhite:
		return []PieceType{
			PieceTypeWhiteKing,
			PieceTypeWhiteQueen,
			PieceTypeWhiteRook,
			PieceTypeWhiteBishop,
			PieceTypeWhiteKnight,
			PieceTypeWhitePawn}
	default:
		return nil
	}
}

// Color returns piece type color.
func (pieceType PieceType) Color() Color {
	switch pieceType {
	case PieceTypeBlackKing, PieceTypeBlackQueen, PieceTypeBlackRook, PieceTypeBlackBishop, PieceTypeBlackKnight,
		PieceTypeBlackPawn:
		return ColorBlack
	case PieceTypeWhiteKing, PieceTypeWhiteQueen, PieceTypeWhiteRook, PieceTypeWhiteBishop, PieceTypeWhiteKnight,
		PieceTypeWhitePawn:
		return ColorWhite
	default:
		return 0
	}
}

// MoveBitboard returns move bitboard for the current piece type and passed origin square.
func (pieceType PieceType) MoveBitboard(origin Square) Bitboard {
	return pieceType.Role().MoveBitboard(origin)
}

// Color returns piece type role.
func (pieceType PieceType) Role() Role {
	switch pieceType {
	case PieceTypeBlackKing, PieceTypeWhiteKing:
		return RoleKing
	case PieceTypeBlackQueen, PieceTypeWhiteQueen:
		return RoleQueen
	case PieceTypeBlackRook, PieceTypeWhiteRook:
		return RoleRook
	case PieceTypeBlackBishop, PieceTypeWhiteBishop:
		return RoleBishop
	case PieceTypeBlackKnight, PieceTypeWhiteKnight:
		return RoleKnight
	case PieceTypeBlackPawn, PieceTypeWhitePawn:
		return RolePawn
	default:
		return 0
	}
}
