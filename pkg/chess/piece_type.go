package chess

import "errors"

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

var pieceTypeFENMap = map[byte]PieceType{
	'k': PieceTypeBlackKing,
	'q': PieceTypeBlackQueen,
	'r': PieceTypeBlackRook,
	'b': PieceTypeBlackBishop,
	'n': PieceTypeBlackKnight,
	'p': PieceTypeBlackPawn,
	'K': PieceTypeWhiteKing,
	'Q': PieceTypeWhiteQueen,
	'R': PieceTypeWhiteRook,
	'B': PieceTypeWhiteBishop,
	'N': PieceTypeWhiteKnight,
	'P': PieceTypeWhitePawn,
}

func NewPieceTypeFromFEN(fen byte) (PieceType, error) {
	pieceType, ok := pieceTypeFENMap[fen]
	if !ok {
		return pieceType, errors.New("unknown byte")
	}
	return pieceType, nil
}
