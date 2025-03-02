package chess

import (
	"errors"
	"fmt"
)

// Piece represents chess pieces.
type Piece uint8

const (
	PieceNil Piece = iota
	PieceWhiteKing
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

var (
	// Mapping of FEN string to corresponding Piece.
	pieceFromFENMap = map[string]Piece{
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

	// Mapping of all piece variants to strings.
	pieceStrings = map[Piece]string{
		PieceNil:         "PieceNil",
		PieceWhiteKing:   "PieceWhiteKing",
		PieceWhiteQueen:  "PieceWhiteQueen",
		PieceWhiteRook:   "PieceWhiteRook",
		PieceWhiteBishop: "PieceWhiteBishop",
		PieceWhiteKnight: "PieceWhiteKnight",
		PieceWhitePawn:   "PieceWhitePawn",
		PieceBlackKing:   "PieceBlackKing",
		PieceBlackQueen:  "PieceBlackQueen",
		PieceBlackRook:   "PieceBlackRook",
		PieceBlackBishop: "PieceBlackBishop",
		PieceBlackKnight: "PieceBlackKnight",
		PieceBlackPawn:   "PieceBlackPawn",
	}
)

// NewPiece returns a piece of passed color and role.
//
// TODO: test.
func NewPiece(color Color, role Role) (Piece, error) {
	switch {
	case color == ColorWhite && role == RoleKing:
		return PieceWhiteKing, nil
	case color == ColorWhite && role == RoleQueen:
		return PieceWhiteQueen, nil
	case color == ColorWhite && role == RoleRook:
		return PieceWhiteRook, nil
	case color == ColorWhite && role == RoleBishop:
		return PieceWhiteBishop, nil
	case color == ColorWhite && role == RoleKnight:
		return PieceWhiteKnight, nil
	case color == ColorWhite && role == RolePawn:
		return PieceWhitePawn, nil
	case color == ColorBlack && role == RoleKing:
		return PieceBlackKing, nil
	case color == ColorBlack && role == RoleQueen:
		return PieceBlackQueen, nil
	case color == ColorBlack && role == RoleRook:
		return PieceBlackRook, nil
	case color == ColorBlack && role == RoleBishop:
		return PieceBlackBishop, nil
	case color == ColorBlack && role == RoleKnight:
		return PieceBlackKnight, nil
	case color == ColorBlack && role == RolePawn:
		return PieceBlackPawn, nil
	default:
		return PieceNil, errors.New("unknown color and role pair")
	}
}

// NewPiecesOfColor returns all pieces of passed color.
//
// TODO: remove allocation.
func NewPiecesOfColor(color Color) ([]Piece, error) {
	switch color {
	case ColorBlack:
		return []Piece{
			PieceBlackKing, PieceBlackQueen, PieceBlackRook, PieceBlackBishop, PieceBlackKnight, PieceBlackPawn}, nil
	case ColorWhite:
		return []Piece{
			PieceWhiteKing, PieceWhiteQueen, PieceWhiteRook, PieceWhiteBishop, PieceWhiteKnight, PieceWhitePawn}, nil
	case ColorNil:
		return nil, errors.New("no pieces")
	default:
		return nil, errors.New("unknown color")
	}
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

// Color returns piece color.
func (piece Piece) Color() (Color, error) {
	switch piece {
	case PieceBlackKing, PieceBlackQueen, PieceBlackRook, PieceBlackBishop, PieceBlackKnight, PieceBlackPawn:
		return ColorBlack, nil
	case PieceWhiteKing, PieceWhiteQueen, PieceWhiteRook, PieceWhiteBishop, PieceWhiteKnight, PieceWhitePawn:
		return ColorWhite, nil
	case PieceNil:
		return ColorNil, errors.New("no color")
	default:
		return ColorNil, errors.New("unknown piece")
	}
}

// NeedPromoInRank returns true if piece need promotion in passed rank.
func (piece Piece) NeedPromoInRank(rank Rank) bool {
	return (piece == PieceWhitePawn && rank == Rank8) || (piece == PieceBlackPawn && rank == Rank1)
}

// IsPawnLongMovePossibleFromRank returns true if current piece is pawn and long move possible from passed rank.
//
// TODO: rewrite without pawn hardcode.
func (piece Piece) IsPawnLongMovePossibleFromRank(rank Rank) bool {
	return (piece == PieceWhitePawn && rank == Rank2) || (piece == PieceBlackPawn && rank == Rank7)
}

// Color returns piece role.
func (piece Piece) Role() (Role, error) {
	switch piece {
	case PieceBlackKing, PieceWhiteKing:
		return RoleKing, nil
	case PieceBlackQueen, PieceWhiteQueen:
		return RoleQueen, nil
	case PieceBlackRook, PieceWhiteRook:
		return RoleRook, nil
	case PieceBlackBishop, PieceWhiteBishop:
		return RoleBishop, nil
	case PieceBlackKnight, PieceWhiteKnight:
		return RoleKnight, nil
	case PieceBlackPawn, PieceWhitePawn:
		return RolePawn, nil
	case PieceNil:
		return RoleNil, errors.New("no role")
	default:
		return RoleNil, errors.New("unknown piece")
	}
}

// String returns string representation of current piece.
func (piece Piece) String() string {
	str, ok := pieceStrings[piece]
	if !ok {
		return fmt.Sprintf("<unknown Piece=%d>", piece)
	}

	return str
}
