package square

import (
	"slices"
	"testing"
)

func TestNewSquare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		file      File
		rank      Rank
		square    Square
		errString string
	}{
		{FileA, Rank1, SquareA1, ""},
		{FileB, Rank2, SquareB2, ""},
		{FileC, Rank3, SquareC3, ""},
		{FileD, Rank4, SquareD4, ""},
		{FileE, Rank5, SquareE5, ""},
		{FileF, Rank6, SquareF6, ""},
		{FileG, Rank7, SquareG7, ""},
		{FileH, Rank8, SquareH8, ""},
		{FileA, RankNil, SquareNil, "invalid rank"},
		{FileNil, Rank1, SquareNil, "invalid file"},
		{FileB, Rank(123), SquareNil, "invalid rank"},
		{File(123), Rank5, SquareNil, "invalid file"},
	}

	for _, test := range tests {
		t.Run(test.square.String(), func(t *testing.T) {
			t.Parallel()

			square, err := NewSquare(test.rank, test.file)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewSquare(%s, %s) expected error %q but got %q", test.rank, test.file, test.errString, err)
			}

			if square != test.square {
				t.Fatalf("NewSquare(%s, %s) expected %s but got %s", test.rank, test.file, test.square, square)
			}
		})
	}
}

func TestNewSquareEnPassantFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		fen       string
		square    Square
		errString string
	}{
		{"a3", SquareA3, ""},
		{"b3", SquareB3, ""},
		{"c3", SquareC3, ""},
		{"d3", SquareD3, ""},
		{"e3", SquareE3, ""},
		{"f3", SquareF3, ""},
		{"g3", SquareG3, ""},
		{"h3", SquareH3, ""},
		{"a6", SquareA6, ""},
		{"b6", SquareB6, ""},
		{"c6", SquareC6, ""},
		{"d6", SquareD6, ""},
		{"e6", SquareE6, ""},
		{"f6", SquareF6, ""},
		{"g6", SquareG6, ""},
		{"h6", SquareH6, ""},
		{"-", SquareNil, ""},
		{"a1", SquareNil, "invalid En Passant rank 1"},
		{"b2", SquareNil, "invalid En Passant rank 2"},
		{"d4", SquareNil, "invalid En Passant rank 4"},
		{"e5", SquareNil, "invalid En Passant rank 5"},
		{"g7", SquareNil, "invalid En Passant rank 7"},
		{"h8", SquareNil, "invalid En Passant rank 8"},
		{"a9", SquareNil, "NewSquareFromFEN(\"a9\"): unknown FEN"},
		{"xyz", SquareNil, "NewSquareFromFEN(\"xyz\"): unknown FEN"},
		{"abc", SquareNil, "NewSquareFromFEN(\"abc\"): unknown FEN"},
	}

	for _, test := range tests {
		t.Run(test.fen, func(t *testing.T) {
			t.Parallel()

			square, err := NewSquareEnPassantFromFEN(test.fen)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewSquareEnPassantFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if square != test.square {
				t.Fatalf("NewSquareEnPassantFromFEN(%q) expected %s but got %s", test.fen, test.square, square)
			}
		})
	}
}

func TestNewSquareFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		fen       string
		square    Square
		errString string
	}{
		{"a1", SquareA1, ""},
		{"b2", SquareB2, ""},
		{"c3", SquareC3, ""},
		{"d4", SquareD4, ""},
		{"e5", SquareE5, ""},
		{"f6", SquareF6, ""},
		{"g7", SquareG7, ""},
		{"h8", SquareH8, ""},
		{"a9", 0, "unknown FEN"},
		{"-", 0, "unknown FEN"},
		{"xyz", 0, "unknown FEN"},
		{"abc", 0, "unknown FEN"},
	}

	for _, test := range tests {
		t.Run(test.fen, func(t *testing.T) {
			t.Parallel()

			square, err := NewSquareFromFEN(test.fen)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewSquareFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if square != test.square {
				t.Fatalf("NewSquareFromFEN(%q) expected %s but got %s", test.fen, test.square, square)
			}
		})
	}
}

func TestSquareFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		square    Square
		file      File
		errString string
	}{
		{SquareA4, FileA, ""},
		{SquareB1, FileB, ""},
		{SquareB2, FileB, ""},
		{SquareC7, FileC, ""},
		{SquareD2, FileD, ""},
		{SquareG3, FileG, ""},
		{SquareH5, FileH, ""},
		{SquareH7, FileH, ""},
		{SquareNil, FileNil, "no file"},
		{Square(65), FileNil, "unknown square"},
		{Square(123), FileNil, "unknown square"},
	}

	for _, test := range tests {
		t.Run(test.square.String(), func(t *testing.T) {
			t.Parallel()

			file, err := test.square.File()
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("%s.File() expected error %q but got %q", test.square, test.errString, err)
			}

			if file != test.file {
				t.Fatalf("%s.File() expected %s but got %s", test.square, test.file, file)
			}
		})
	}
}

func TestSquareRank(t *testing.T) {
	t.Parallel()

	tests := []struct {
		square    Square
		rank      Rank
		errString string
	}{
		{SquareA1, Rank1, ""},
		{SquareH1, Rank1, ""},
		{SquareB2, Rank2, ""},
		{SquareC3, Rank3, ""},
		{SquareD4, Rank4, ""},
		{SquareE5, Rank5, ""},
		{SquareH6, Rank6, ""},
		{SquareF7, Rank7, ""},
		{SquareA8, Rank8, ""},
		{SquareH8, Rank8, ""},
		{SquareNil, RankNil, "no rank"},
		{Square(65), RankNil, "unknown square"},
		{Square(222), RankNil, "unknown square"},
	}

	for _, test := range tests {
		t.Run(test.square.String(), func(t *testing.T) {
			t.Parallel()

			rank, err := test.square.Rank()
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("%s.Rank() expected error %q but got %q", test.square, test.errString, err)
			}

			if rank != test.rank {
				t.Fatalf("%s.Rank() expected %s but got %s", test.square, test.rank, rank)
			}
		})
	}
}

func TestSquareString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		square Square
		str    string
	}{
		{SquareNil, "SquareNil"},
		{SquareA1, "SquareA1"},
		{SquareA8, "SquareA8"},
		{SquareD1, "SquareD1"},
		{SquareE8, "SquareE8"},
		{SquareH1, "SquareH1"},
		{SquareH8, "SquareH8"},
		{Square(123), "<unknown Square=123>"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			t.Parallel()

			str := test.square.String()
			if str != test.str {
				t.Fatalf("%s.String() expected %q but got %q", test.str, test.str, str)
			}
		})
	}
}
