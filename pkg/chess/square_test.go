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
