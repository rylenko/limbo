package chess

import (
	"reflect"
	"testing"
)

func TestNewPositionFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		fen       string
		position  *Position
		errString string
	}{
		{
			"start",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			NewPosition(
				NewBoard(map[PieceType]Bitboard{
					PieceTypeWhiteKing:   Bitboard(0x0800000000000000),
					PieceTypeWhiteQueen:  Bitboard(0x1000000000000000),
					PieceTypeWhiteRook:   Bitboard(0x8100000000000000),
					PieceTypeWhiteBishop: Bitboard(0x2400000000000000),
					PieceTypeWhiteKnight: Bitboard(0x4200000000000000),
					PieceTypeWhitePawn:   Bitboard(0x00FF000000000000),
					PieceTypeBlackKing:   Bitboard(0x0000000000000008),
					PieceTypeBlackQueen:  Bitboard(0x0000000000000010),
					PieceTypeBlackRook:   Bitboard(0x0000000000000081),
					PieceTypeBlackBishop: Bitboard(0x0000000000000024),
					PieceTypeBlackKnight: Bitboard(0x0000000000000042),
					PieceTypeBlackPawn:   Bitboard(0x000000000000FF00),
				}),
				ColorWhite,
				NewCastlingRights(ColorSideWhiteKing, ColorSideWhiteQueen, ColorSideBlackKing, ColorSideBlackQueen),
				nil,
				0,
				1,
			),
			"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			position, err := NewPositionFromFEN(test.fen)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewPositionFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if !reflect.DeepEqual(position, test.position) {
				t.Fatalf("NewPositionFromFEN(%q) expected %+v but got %+v", test.fen, test.position, position)
			}
		})
	}
}
