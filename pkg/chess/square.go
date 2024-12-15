package chess

import (
	"errors"
	"fmt"
	"slices"
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

var (
	squareFromFENMap = map[string]Square{
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
)

// NewSquare creates a new square of passed file and rank.
func NewSquare(rank Rank, file File) (Square, error) {
	rankSquares, err := NewSquaresOfRank(rank)
	if err != nil {
		return SquareNil, fmt.Errorf("NewSquaresOfRank(%s): %w", rank, err)
	}

	fileSquares, err := NewSquaresOfFile(file)
	if err != nil {
		return SquareNil, fmt.Errorf("NewSquaresOfFile(%s): %w", file, err)
	}

	for _, fileSquare := range fileSquares {
		if slices.Contains(rankSquares, fileSquare) {
			return fileSquare, nil
		}
	}

	return SquareNil, errors.New("unreachable statement")
}

// NewSquaresOfFile creates a slice of squares of passed file.
func NewSquaresOfFile(file File) ([]Square, error) {
	switch file {
	case FileA:
		return []Square{SquareA1, SquareA2, SquareA3, SquareA4, SquareA5, SquareA6, SquareA7, SquareA8}, nil
	case FileB:
		return []Square{SquareB1, SquareB2, SquareB3, SquareB4, SquareB5, SquareB6, SquareB7, SquareB8}, nil
	case FileC:
		return []Square{SquareC1, SquareC2, SquareC3, SquareC4, SquareC5, SquareC6, SquareC7, SquareC8}, nil
	case FileD:
		return []Square{SquareD1, SquareD2, SquareD3, SquareD4, SquareD5, SquareD6, SquareD7, SquareD8}, nil
	case FileE:
		return []Square{SquareE1, SquareE2, SquareE3, SquareE4, SquareE5, SquareE6, SquareE7, SquareE8}, nil
	case FileF:
		return []Square{SquareF1, SquareF2, SquareF3, SquareF4, SquareF5, SquareF6, SquareF7, SquareF8}, nil
	case FileG:
		return []Square{SquareG1, SquareG2, SquareG3, SquareG4, SquareG5, SquareG6, SquareG7, SquareG8}, nil
	case FileH:
		return []Square{SquareH1, SquareH2, SquareH3, SquareH4, SquareH5, SquareH6, SquareH7, SquareH8}, nil
	case FileNil:
		return nil, errors.New("no squares")
	default:
		return nil, errors.New("unknown file")
	}
}

// NewSquaresOfRank creates a slice of squares of passed rank.
func NewSquaresOfRank(rank Rank) ([]Square, error) {
	switch rank {
	case Rank1:
		return []Square{SquareA1, SquareB1, SquareC1, SquareD1, SquareE1, SquareF1, SquareG1, SquareH1}, nil
	case Rank2:
		return []Square{SquareA2, SquareB2, SquareC2, SquareD2, SquareE2, SquareF2, SquareG2, SquareH2}, nil
	case Rank3:
		return []Square{SquareA3, SquareB3, SquareC3, SquareD3, SquareE3, SquareF3, SquareG3, SquareH3}, nil
	case Rank4:
		return []Square{SquareA4, SquareB4, SquareC4, SquareD4, SquareE4, SquareF4, SquareG4, SquareH4}, nil
	case Rank5:
		return []Square{SquareA5, SquareB5, SquareC5, SquareD5, SquareE5, SquareF5, SquareG5, SquareH5}, nil
	case Rank6:
		return []Square{SquareA6, SquareB6, SquareC6, SquareD6, SquareE6, SquareF6, SquareG6, SquareH6}, nil
	case Rank7:
		return []Square{SquareA7, SquareB7, SquareC7, SquareD7, SquareE7, SquareF7, SquareG7, SquareH7}, nil
	case Rank8:
		return []Square{SquareA8, SquareB8, SquareC8, SquareD8, SquareE8, SquareF8, SquareG8, SquareH8}, nil
	case RankNil:
		return nil, errors.New("no squares")
	default:
		return nil, errors.New("unknown rank")
	}
}

// NewSquareEnPassantFromFEN parses En Passant FEN to corresponding Square or returns an error.
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

	rank, err := square.Rank()
	if err != nil {
		return SquareNil, fmt.Errorf("%s.Rank(): %w", square, err)
	}

	if !RolePawn.IsEnPassantPossibleInRank(rank) {
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
	switch square {
	case SquareA1, SquareA2, SquareA3, SquareA4, SquareA5, SquareA6, SquareA7, SquareA8:
		return FileA, nil
	case SquareB1, SquareB2, SquareB3, SquareB4, SquareB5, SquareB6, SquareB7, SquareB8:
		return FileB, nil
	case SquareC1, SquareC2, SquareC3, SquareC4, SquareC5, SquareC6, SquareC7, SquareC8:
		return FileC, nil
	case SquareD1, SquareD2, SquareD3, SquareD4, SquareD5, SquareD6, SquareD7, SquareD8:
		return FileD, nil
	case SquareE1, SquareE2, SquareE3, SquareE4, SquareE5, SquareE6, SquareE7, SquareE8:
		return FileE, nil
	case SquareF1, SquareF2, SquareF3, SquareF4, SquareF5, SquareF6, SquareF7, SquareF8:
		return FileF, nil
	case SquareG1, SquareG2, SquareG3, SquareG4, SquareG5, SquareG6, SquareG7, SquareG8:
		return FileG, nil
	case SquareH1, SquareH2, SquareH3, SquareH4, SquareH5, SquareH6, SquareH7, SquareH8:
		return FileH, nil
	case SquareNil:
		return FileNil, errors.New("no file")
	default:
		return FileNil, errors.New("unknown square")
	}
}

// Rank returns rank of the current square.
func (square Square) Rank() (Rank, error) {
	switch square {
	case SquareA1, SquareB1, SquareC1, SquareD1, SquareE1, SquareF1, SquareG1, SquareH1:
		return Rank1, nil
	case SquareA2, SquareB2, SquareC2, SquareD2, SquareE2, SquareF2, SquareG2, SquareH2:
		return Rank2, nil
	case SquareA3, SquareB3, SquareC3, SquareD3, SquareE3, SquareF3, SquareG3, SquareH3:
		return Rank3, nil
	case SquareA4, SquareB4, SquareC4, SquareD4, SquareE4, SquareF4, SquareG4, SquareH4:
		return Rank4, nil
	case SquareA5, SquareB5, SquareC5, SquareD5, SquareE5, SquareF5, SquareG5, SquareH5:
		return Rank5, nil
	case SquareA6, SquareB6, SquareC6, SquareD6, SquareE6, SquareF6, SquareG6, SquareH6:
		return Rank6, nil
	case SquareA7, SquareB7, SquareC7, SquareD7, SquareE7, SquareF7, SquareG7, SquareH7:
		return Rank7, nil
	case SquareA8, SquareB8, SquareC8, SquareD8, SquareE8, SquareF8, SquareG8, SquareH8:
		return Rank8, nil
	case SquareNil:
		return RankNil, errors.New("no rank")
	default:
		return RankNil, errors.New("unknown square")
	}
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
