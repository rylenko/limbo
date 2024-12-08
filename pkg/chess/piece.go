package chess

import "errors"

// Piece represents chess pieces.
type Piece uint8

const (
	PieceWhiteKing Piece = iota
	PieceWhiteQueen
	PieceWhiteRook
	PieceWhiteBishop
	PieceWhiteKnight
	PieceWhitePawn
	PieceBlackKing
	PieceBlackQueen
	PieceBlackRook
	PieceBlackBishop
	PieceBlackKnight
	PieceBlackPawn
)

// Mapping of FEN string to corresponding Piece.
var pieceFromFENMap = map[string]Piece{
	"k": PieceBlackKing,
	"q": PieceBlackQueen,
	"r": PieceBlackRook,
	"b": PieceBlackBishop,
	"n": PieceBlackKnight,
	"p": PieceBlackPawn,
	"K": PieceWhiteKing,
	"Q": PieceWhiteQueen,
	"R": PieceWhiteRook,
	"B": PieceWhiteBishop,
	"N": PieceWhiteKnight,
	"P": PieceWhitePawn,
}

// NewPieceFromFEN parses FEN to corresponding Piece or returns an error.
//
// FEN argument examples: "k", "q", "r", "b", "n", "p" for black pieces and and the same, but in upper case, for whites.
func NewPieceFromFEN(fen string) (Piece, error) {
	piece, ok := pieceFromFENMap[fen]
	if !ok {
		return piece, errors.New("unknown FEN")
	}
	return piece, nil
}

// NewPiecesFromColor returns slice of pieces of passed Color.
func NewPiecesOfColor(color Color) []Piece {
	switch color {
	case ColorBlack:
		return []Piece{PieceBlackKing, PieceBlackQueen, PieceBlackRook, PieceBlackBishop, PieceBlackKnight, PieceBlackPawn}
	case ColorWhite:
		return []Piece{PieceWhiteKing, PieceWhiteQueen, PieceWhiteRook, PieceWhiteBishop, PieceWhiteKnight, PieceWhitePawn}
	default:
		return nil
	}
}

// Color returns piece color.
func (piece Piece) Color() Color {
	switch piece {
	case PieceBlackKing, PieceBlackQueen, PieceBlackRook, PieceBlackBishop, PieceBlackKnight, PieceBlackPawn:
		return ColorBlack
	case PieceWhiteKing, PieceWhiteQueen, PieceWhiteRook, PieceWhiteBishop, PieceWhiteKnight, PieceWhitePawn:
		return ColorWhite
	default:
		return 0
	}
}

// Color returns piece role.
func (piece Piece) Role() Role {
	switch piece {
	case PieceBlackKing, PieceWhiteKing:
		return RoleKing
	case PieceBlackQueen, PieceWhiteQueen:
		return RoleQueen
	case PieceBlackRook, PieceWhiteRook:
		return RoleRook
	case PieceBlackBishop, PieceWhiteBishop:
		return RoleBishop
	case PieceBlackKnight, PieceWhiteKnight:
		return RoleKnight
	case PieceBlackPawn, PieceWhitePawn:
		return RolePawn
	default:
		return 0
	}
}
