package chess

import "errors"

type PieceType uint8

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

var (
	PieceTypes = []PieceType{
		PieceTypeBlackKing, PieceTypeBlackQueen, PieceTypeBlackRook, PieceTypeBlackBishop, PieceTypeBlackKnight,
		PieceTypeBlackPawn,

		PieceTypeWhiteKing, PieceTypeWhiteQueen, PieceTypeWhiteRook, PieceTypeWhiteBishop, PieceTypeWhiteKnight,
		PieceTypeWhitePawn,
	}
	pieceTypeFENMap = map[byte]PieceType{
		byte('k'): PieceTypeBlackKing,
		byte('q'): PieceTypeBlackQueen,
		byte('r'): PieceTypeBlackRook,
		byte('b'): PieceTypeBlackBishop,
		byte('n'): PieceTypeBlackKnight,
		byte('p'): PieceTypeBlackPawn,
		byte('K'): PieceTypeWhiteKing,
		byte('Q'): PieceTypeWhiteQueen,
		byte('R'): PieceTypeWhiteRook,
		byte('B'): PieceTypeWhiteBishop,
		byte('N'): PieceTypeWhiteKnight,
		byte('P'): PieceTypeWhitePawn,
	}
)

func NewPieceTypeFromFEN(fen byte) (PieceType, error) {
	pieceType, ok := pieceTypeFENMap[fen]
	if !ok {
		return pieceType, errors.New("unknown FEN byte")
	}
	return pieceType, nil
}
