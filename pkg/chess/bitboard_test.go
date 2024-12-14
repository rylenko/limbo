package chess

import (
	"slices"
	"testing"
)

func TestBitboardSetSquaresAndGetSquares(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		squares      []Square
		bitboard     Bitboard
		setErrString string
	}{
		{
			"diagonal",
			[]Square{SquareA1, SquareB2, SquareC3, SquareD4, SquareE5, SquareF6, SquareG7, SquareH8},
			0x8040201008040201,
			"",
		},
		{
			"pawns",
			[]Square{
				SquareA2, SquareB2, SquareC2, SquareD2, SquareE2, SquareF2, SquareG2, SquareH2,
				SquareA7, SquareB7, SquareC7, SquareD7, SquareE7, SquareF7, SquareG7, SquareH7,
			},
			0x00FF00000000FF00,
			"",
		},
		{"left-middle_right-middle", []Square{SquareA4, SquareH4, SquareA5, SquareH5}, 0x0000008181000000, ""},
		{
			"invalid set",
			[]Square{SquareA1, SquareB1, Square(123)},
			0xC000000000000000,
			"invalid square <unknown Square=123>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			bitboard, err := BitboardNil.SetSquares(test.squares...)
			if (err == nil && test.setErrString != "") || (err != nil && err.Error() != test.setErrString) {
				t.Fatalf("SetSquares(%v) expected error %q but got %q", test.squares, test.setErrString, err)
			}

			if bitboard != test.bitboard {
				t.Fatalf("SetSquares(%v) expected 0x%X but got 0x%X", test.squares, test.bitboard, bitboard)
			}

			squares := bitboard.GetSquares()
			if !slices.Equal(squares, test.squares) {
				t.Fatalf("GetSquares(0x%X) expected %v but got %v", test.bitboard, test.squares, squares)
			}
		})
	}
}
