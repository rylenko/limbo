package chess

import (
	"slices"
	"testing"
)

func TestBitboardGetSquares(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		bitboard Bitboard
		squares  []Square
	}{
		{
			"diagonal",
			Bitboard(0x8040201008040201),
			[]Square{SquareA1, SquareB2, SquareC3, SquareD4, SquareE5, SquareF6, SquareG7, SquareH8},
		},
		{
			"pawns",
			Bitboard(0x00FF00000000FF00),
			[]Square{
				SquareA2, SquareB2, SquareC2, SquareD2, SquareE2, SquareF2, SquareG2, SquareH2,
				SquareA7, SquareB7, SquareC7, SquareD7, SquareE7, SquareF7, SquareG7, SquareH7,
			},
		},
		{"left-middle_right-middle", Bitboard(0x0000008181000000), []Square{SquareA4, SquareH4, SquareA5, SquareH5}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			squares := test.bitboard.GetSquares()
			if !slices.Equal(squares, test.squares) {
				t.Fatalf("GetSquares() expected %v but got %v", test.squares, squares)
			}
		})
	}
}

func TestBitboardSetSquares(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		squares  []Square
		bitboard Bitboard
	}{
		{
			"diagonal",
			[]Square{SquareA1, SquareB2, SquareC3, SquareD4, SquareE5, SquareF6, SquareG7, SquareH8},
			Bitboard(0x8040201008040201),
		},
		{
			"pawns",
			[]Square{
				SquareA2, SquareB2, SquareC2, SquareD2, SquareE2, SquareF2, SquareG2, SquareH2,
				SquareA7, SquareB7, SquareC7, SquareD7, SquareE7, SquareF7, SquareG7, SquareH7,
			},
			Bitboard(0x00FF00000000FF00),
		},
		{"left-middle_right-middle", []Square{SquareA4, SquareH4, SquareA5, SquareH5}, Bitboard(0x0000008181000000)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var bitboard Bitboard
			bitboard = bitboard.SetSquares(test.squares...)

			if bitboard != test.bitboard {
				t.Fatalf("Bitboard of squares %v expected 0x%X but got 0x%X", test.squares, test.bitboard, bitboard)
			}
		})
	}
}
