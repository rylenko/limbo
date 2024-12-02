package chess

import "errors"

// PieceType represents chess piece types.
type PieceType int8

const (
	PieceTypeBlackKing PieceType = iota
	PieceTypeBlackQueen
	PieceTypeBlackRook
	PieceTypeBlackBishop
	PieceTypeBlackKnight
	PieceTypeBlackPawn
	PieceTypeWhiteKing
	PieceTypeWhiteQueen
	PieceTypeWhiteRook
	PieceTypeWhiteBishop
	PieceTypeWhiteKnight
	PieceTypeWhitePawn
)

// Mapping of FEN byte to corresponding PieceType.
var pieceTypeFENMap = map[string]PieceType{
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
	pieceType, ok := pieceTypeFENMap[fen]
	if !ok {
		return pieceType, errors.New("unknown FEN")
	}
	return pieceType, nil
}
