package chess

import "testing"

func TestNewSquare(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		file   File
		rank   Rank
		square Square
	}{
		{"A1", FileA, Rank1, SquareA1},
		{"B2", FileB, Rank2, SquareB2},
		{"C3", FileC, Rank3, SquareC3},
		{"D4", FileD, Rank4, SquareD4},
		{"E5", FileE, Rank5, SquareE5},
		{"F6", FileF, Rank6, SquareF6},
		{"G7", FileG, Rank7, SquareG7},
		{"H8", FileH, Rank8, SquareH8},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			square := NewSquare(test.file, test.rank)

			if square != test.square {
				t.Fatalf("NewSquare(%d, %d) expected %d but got %d", test.file, test.rank, test.square, square)
			}
		})
	}
}

func TestNewSquareEnPassantFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		fen       string
		square    *Square
		errString string
	}{
		{"a3", Ptr(SquareA3), ""},
		{"b3", Ptr(SquareB3), ""},
		{"c3", Ptr(SquareC3), ""},
		{"d3", Ptr(SquareD3), ""},
		{"e3", Ptr(SquareE3), ""},
		{"f3", Ptr(SquareF3), ""},
		{"g3", Ptr(SquareG3), ""},
		{"h3", Ptr(SquareH3), ""},
		{"a6", Ptr(SquareA6), ""},
		{"b6", Ptr(SquareB6), ""},
		{"c6", Ptr(SquareC6), ""},
		{"d6", Ptr(SquareD6), ""},
		{"e6", Ptr(SquareE6), ""},
		{"f6", Ptr(SquareF6), ""},
		{"g6", Ptr(SquareG6), ""},
		{"h6", Ptr(SquareH6), ""},
		{"-", nil, ""},
		{"a1", nil, "invalid rank 0"},
		{"b2", nil, "invalid rank 1"},
		{"d4", nil, "invalid rank 3"},
		{"e5", nil, "invalid rank 4"},
		{"g7", nil, "invalid rank 6"},
		{"h8", nil, "invalid rank 7"},
		{"a9", nil, "NewSquareFromFEN(\"a9\"): unknown FEN"},
		{"xyz", nil, "NewSquareFromFEN(\"xyz\"): unknown FEN"},
		{"abc", nil, "NewSquareFromFEN(\"abc\"): unknown FEN"},
	}

	for _, test := range tests {
		t.Run(test.fen, func(t *testing.T) {
			t.Parallel()

			square, err := NewSquareEnPassantFromFEN(test.fen)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewSquareEnPassantFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if (square != nil && test.square == nil) ||
				(square == nil && test.square != nil) ||
				(square != nil && test.square != nil && *square != *test.square) {
				t.Fatalf("NewSquareEnPassantFromFEN(%q) expected %d but got %d", test.fen, test.square, square)
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
				t.Fatalf("NewSquareFromFEN(%q) expected %d but got %d", test.fen, test.square, square)
			}
		})
	}
}

func TestSquareFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		square Square
		file   File
	}{
		{"B1", SquareB1, FileB},
		{"C7", SquareC7, FileC},
		{"D2", SquareD2, FileD},
		{"A4", SquareA4, FileA},
		{"G3", SquareG3, FileG},
		{"H5", SquareH5, FileH},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if file := test.square.File(); file != test.file {
				t.Fatalf("Square(%d).File() expected %d but got %d", test.square, test.file, file)
			}
		})
	}
}

func TestSquareOrder(t *testing.T) {
	t.Parallel()

	if SquareA1 >= SquareB1 {
		t.Fatalf("SquareA1=%d >= SquareB1=%d", SquareA1, SquareB1)
	}
}

func TestSquareRank(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		square Square
		rank   Rank
	}{
		{"A1", SquareA1, Rank1},
		{"H1", SquareH1, Rank1},
		{"D4", SquareD4, Rank4},
		{"E5", SquareE5, Rank5},
		{"A8", SquareA8, Rank8},
		{"H8", SquareH8, Rank8},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if rank := test.square.Rank(); rank != test.rank {
				t.Fatalf("Square(%d).Rank() expected %d but got %d", test.square, test.rank, rank)
			}
		})
	}
}
