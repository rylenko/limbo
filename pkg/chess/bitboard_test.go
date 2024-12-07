package chess

import (
	"slices"
	"testing"
)

func TestBitboardSetSquaresAndGetSquares(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		squares  []Square
		bitboard Bitboard
	}{
		{
			"diagonal",
			[]Square{SquareA1, SquareB2, SquareC3, SquareD4, SquareE5, SquareF6, SquareG7, SquareH8},
			0x8040201008040201,
		},
		{
			"pawns",
			[]Square{
				SquareA2, SquareB2, SquareC2, SquareD2, SquareE2, SquareF2, SquareG2, SquareH2,
				SquareA7, SquareB7, SquareC7, SquareD7, SquareE7, SquareF7, SquareG7, SquareH7,
			},
			0x00FF00000000FF00,
		},
		{"left-middle_right-middle", []Square{SquareA4, SquareH4, SquareA5, SquareH5}, 0x0000008181000000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var bitboard Bitboard

			bitboard = bitboard.SetSquares(test.squares...)
			if bitboard != test.bitboard {
				t.Fatalf("Bitboard of squares %v expected 0x%X but got 0x%X", test.squares, test.bitboard, bitboard)
			}

			squares := bitboard.GetSquares()
			if !slices.Equal(squares, test.squares) {
				t.Fatalf("Squares of bitboard 0x%X expected %v but got %v", test.bitboard, test.squares, squares)
			}
		})
	}
}
