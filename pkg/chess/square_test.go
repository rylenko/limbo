package chess

import "testing"

func TestSquareOrder(t *testing.T) {
	t.Parallel()

	a1Number := uint8(SquareA1)
	b1Number := uint8(SquareB1)

	if a1Number >= b1Number {
		t.Fatalf("SquareA1=%d must be less than SquareB1=%d", a1Number, b1Number)
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

			gotSquare := NewSquare(test.file, test.rank)

			if gotSquare != test.square {
				t.Fatalf("Square of file %d and rank %d expected %d but got %d", test.file, test.rank, test.square, gotSquare)
			}
		})
	}
}
