package chess

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	testBoardHarder = NewBoard(map[PieceType]Bitboard{
		PieceTypeWhiteKing:   Bitboard(0x0800000000000000),
		PieceTypeWhiteQueen:  Bitboard(0x0000008000000000),
		PieceTypeWhiteRook:   Bitboard(0x8000000000010000),
		PieceTypeWhiteBishop: Bitboard(0x2000100000000000),
		PieceTypeWhiteKnight: Bitboard(0x4000040000000000),
		PieceTypeWhitePawn:   Bitboard(0x00D6002802000000),
		PieceTypeBlackKing:   Bitboard(0x0000000000000004),
		PieceTypeBlackQueen:  Bitboard(0x0000000000000010),
		PieceTypeBlackRook:   Bitboard(0x0000000000000081),
		PieceTypeBlackBishop: Bitboard(0x0000000000000820),
		PieceTypeBlackKnight: Bitboard(0x0000000000240000),
		PieceTypeBlackPawn:   Bitboard(0x000000004018A500),
	})

	testBoardStart = NewBoard(map[PieceType]Bitboard{
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
	})
)

func TestNewBoardFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		fen       string
		board     *Board
		errString string
	}{
		{"start", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", testBoardStart, ""},
		{"harder", "r1bq1k1r/p1p1bp1p/2nppn1R/1p4P1/Q1P1P3/3B1N2/PP1P1PP1/RNB1K3", testBoardHarder, ""},
		{
			"not enough parts",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP",
			nil,
			fmt.Sprintf("required %d parts but got 7", ranksCount)},
		{
			"too many parts",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR/extra-part",
			nil,
			fmt.Sprintf("required %d parts but got 9", ranksCount),
		},
		{
			"invalid piece type FEN",
			"rnbqkbnr/pppXpppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			"part #1, byte #3, NewPieceTypeFromFEN(\"X\"): unknown FEN",
		},
		{
			"not enough files",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPP/RNBQKBNR",
			nil,
			fmt.Sprintf("required %d files but got 7 in part #6", filesCount),
		},
		{
			"too many files",
			"rrnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			fmt.Sprintf("required %d files but got 9 in part #0", filesCount),
		},
		{
			"not enough offsets",
			"rnbqkbnr/pppppppp/8/8/6/8/PPPPPPPP/RNBQKBNR",
			nil,
			fmt.Sprintf("required %d files but got 6 in part #4", filesCount),
		},
		{
			"too many offsets",
			"rnbqkbnr/pppppppp/9/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			fmt.Sprintf("required %d files but got 9 in part #2", filesCount),
		},
		{
			"not enough offsets and pieces",
			"r2bnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			fmt.Sprintf("required %d files but got 6 in part #0", filesCount),
		},
		{
			"too many offsets and pieces",
			"rnbqkbnr/p6ppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			fmt.Sprintf("required %d files but got 10 in part #1", filesCount),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			board, err := NewBoardFromFEN(test.fen)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewBoardFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if !reflect.DeepEqual(board, test.board) {
				t.Fatalf("NewBoardFromFEN(%q) expected %+v but got %+v", test.fen, test.board, board)
			}
		})
	}
}

func TestBoardColorBitboard(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		board          *Board
		colorBitboards map[Color]Bitboard
	}{

		{"start", testBoardStart, map[Color]Bitboard{
			ColorWhite: Bitboard(0xFFFF000000000000),
			ColorBlack: Bitboard(0x000000000000FFFF),
		}},
		{"harder", testBoardHarder, map[Color]Bitboard{
			ColorWhite: Bitboard(0xE8D614A802010000),
			ColorBlack: Bitboard(0x00000000403CADB5),
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for color, expectedBitboard := range test.colorBitboards {
				gotBitboard := test.board.ColorBitboard(color)
				if gotBitboard != expectedBitboard {
					t.Fatalf("ColorBitboard(%d) expected bitboard 0x%X but got 0x%X", color, expectedBitboard, gotBitboard)
				}
			}
		})
	}
}

func TestBoardUnoccupiedBitboard(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		board    *Board
		bitboard Bitboard
	}{
		{"start", testBoardStart, Bitboard(0x0000FFFFFFFF0000)},
		{"harder", testBoardHarder, Bitboard(0x1729EB57BDC2524A)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			bitboard := test.board.UnoccupiedBitboard()
			if bitboard != test.bitboard {
				t.Fatalf("UnoccupiedBitboard() expected bitboard 0x%X but got 0x%X", test.bitboard, bitboard)
			}
		})
	}
}
