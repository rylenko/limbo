package chess

import "testing"

func TestSquareOrder(t *testing.T) {
	t.Parallel()

	if SquareA1 >= SquareB1 {
		t.Fatalf("SquareA1=%d >= SquareB1=%d", SquareA1, SquareB1)
	}
}

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
