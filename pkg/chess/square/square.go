package square

import (
	"errors"
	"fmt"
	"slices"

	"github.com/rylenko/limbo/pkg/chess/piece"
)

// Square represents square on the chess board.
type Square uint8

const (
	SquareNil Square = iota
	SquareA1
	SquareB1
	SquareC1
	SquareD1
	SquareE1
	SquareF1
	SquareG1
	SquareH1
	SquareA2
	SquareB2
	SquareC2
	SquareD2
	SquareE2
	SquareF2
	SquareG2
	SquareH2
	SquareA3
	SquareB3
	SquareC3
	SquareD3
	SquareE3
	SquareF3
	SquareG3
	SquareH3
	SquareA4
	SquareB4
	SquareC4
	SquareD4
	SquareE4
	SquareF4
	SquareG4
	SquareH4
	SquareA5
	SquareB5
	SquareC5
	SquareD5
	SquareE5
	SquareF5
	SquareG5
	SquareH5
	SquareA6
	SquareB6
	SquareC6
	SquareD6
	SquareE6
	SquareF6
	SquareG6
	SquareH6
	SquareA7
	SquareB7
	SquareC7
	SquareD7
	SquareE7
	SquareF7
	SquareG7
	SquareH7
	SquareA8
	SquareB8
	SquareC8
	SquareD8
	SquareE8
	SquareF8
	SquareG8
	SquareH8
)

var squareFromFENMap = map[string]Square{
	"a1": SquareA1,
	"b1": SquareB1,
	"c1": SquareC1,
	"d1": SquareD1,
	"e1": SquareE1,
	"f1": SquareF1,
	"g1": SquareG1,
	"h1": SquareH1,
	"a2": SquareA2,
	"b2": SquareB2,
	"c2": SquareC2,
	"d2": SquareD2,
	"e2": SquareE2,
	"f2": SquareF2,
	"g2": SquareG2,
	"h2": SquareH2,
	"a3": SquareA3,
	"b3": SquareB3,
	"c3": SquareC3,
	"d3": SquareD3,
	"e3": SquareE3,
	"f3": SquareF3,
	"g3": SquareG3,
	"h3": SquareH3,
	"a4": SquareA4,
	"b4": SquareB4,
	"c4": SquareC4,
	"d4": SquareD4,
	"e4": SquareE4,
	"f4": SquareF4,
	"g4": SquareG4,
	"h4": SquareH4,
	"a5": SquareA5,
	"b5": SquareB5,
	"c5": SquareC5,
	"d5": SquareD5,
	"e5": SquareE5,
	"f5": SquareF5,
	"g5": SquareG5,
	"h5": SquareH5,
	"a6": SquareA6,
	"b6": SquareB6,
	"c6": SquareC6,
	"d6": SquareD6,
	"e6": SquareE6,
	"f6": SquareF6,
	"g6": SquareG6,
	"h6": SquareH6,
	"a7": SquareA7,
	"b7": SquareB7,
	"c7": SquareC7,
	"d7": SquareD7,
	"e7": SquareE7,
	"f7": SquareF7,
	"g7": SquareG7,
	"h7": SquareH7,
	"a8": SquareA8,
	"b8": SquareB8,
	"c8": SquareC8,
	"d8": SquareD8,
	"e8": SquareE8,
	"f8": SquareF8,
	"g8": SquareG8,
	"h8": SquareH8,
}

// NewSquare creates a new square using passed file and rank.
func NewSquare(rank Rank, file File) (Square, error) {
	if rank == RankNil || rank > Rank8 {
		return SquareNil, errors.New("invalid rank")
	}

	if file == FileNil || file > FileH {
		return SquareNil, errors.New("invalid file")
	}

	return Square((rank - 1) * len(Ranks) + file)
}

// NewSquareEnPassantFromFEN parses en passant FEN to corresponding Square or returns an error.
//
// FEN argument examples: "-", "a3", "h6".
func NewSquareEnPassantFromFEN(fen string) (Square, error) {
	if fen == "-" {
		return SquareNil, nil
	}

	square, err := NewSquareFromFEN(fen)
	if err != nil {
		return SquareNil, fmt.Errorf("NewSquareFromFEN(%q): %w", fen, err)
	}

	if square.unsafeRank().EnPassantIsPossible() {
		return SquareNil, fmt.Errorf("invalid En Passant rank %d", rank)
	}

	return square, nil
}

// NewSquareFromFEN parses square FEN to correspoding Square or returns an error.
//
// FEN argument examples: "a1", "h8".
func NewSquareFromFEN(fen string) (Square, error) {
	square, ok := squareFromFENMap[fen]
	if !ok {
		return square, errors.New("unknown FEN")
	}

	return square, nil
}

// File returns file of the current square.
func (square Square) File() (File, error) {
	if square == SquareNil || square > SquareH8 {
		return RankNil, errors.New("unknown square")
	}

	return square.unsafeFile(), nil
}

// Rank returns rank of the current square.
func (square Square) Rank() (Rank, error) {
	if square == SquareNil || square > SquareH8 {
		return RankNil, errors.New("unknown square")
	}

	return square.unsafeRank(), nil
}

// String returns string represetntation of current square.
func (square Square) String() string {
	if square == SquareNil {
		return "SquareNil"
	}

	rank, err := square.Rank()
	if err != nil {
		return fmt.Sprintf("<unknown Square=%d>", square)
	}
	rankDigit := uint8('1') + (uint8(rank) - 1)

	file, err := square.File()
	if err != nil {
		return fmt.Sprintf("<unknown Square=%d>", square)
	}
	fileLetter := uint8('A') + (uint8(file) - 1)

	str := fmt.Sprintf("Square%c%c", fileLetter, rankDigit)

	return str
}

// unsafeFile returns file using unvalidated calculations.
func (square Square) unsafeFile() Rank {
	return Rank(square % len(Files) + 1)
}

// unsafeRank returns rank using unvalidated calculations.
func (square Square) unsafeRank() Rank {
	return Rank(square / len(Ranks) + 1)
}
