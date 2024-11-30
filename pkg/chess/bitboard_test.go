package chess

import (
	"testing"

	set "github.com/deckarep/golang-set/v2"
)

func TestNewBitboard(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		squares  set.Set[Square]
		bitboard bitboard
	}{
		{
			"diagonal",
			set.NewThreadUnsafeSet(SquareA1, SquareB2, SquareC3, SquareD4, SquareE5, SquareF6, SquareG7, SquareH8),
			bitboard(0x8040201008040201),
		},
		{
			"left-middle_right-middle",
			set.NewThreadUnsafeSet(SquareA4, SquareA5, SquareH4, SquareH5),
			bitboard(0x8181000000),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			gotBitboard := newBitboard(test.squares)

			if gotBitboard != test.bitboard {
				t.Fatalf("Bitboard of squares %v expected %d but got %d", test.squares, test.bitboard, gotBitboard)
			}
		})
	}
}
