package chess

import "testing"

func TestBitboardSet(t *testing.T) {
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
		{"left-middle_right-middle", []Square{SquareA4, SquareA5, SquareH4, SquareH5}, Bitboard(0x0000008181000000)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var bitboard Bitboard
			for _, square := range test.squares {
				bitboard = bitboard.Set(square)
			}

			if bitboard != test.bitboard {
				t.Fatalf("Bitboard of squares %v expected %d but got %d", test.squares, test.bitboard, bitboard)
			}
		})
	}
}
