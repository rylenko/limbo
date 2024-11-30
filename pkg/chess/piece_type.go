package chess

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
)
