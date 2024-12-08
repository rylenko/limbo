package chess

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	testBoardHarder = NewBoard(map[Piece]Bitboard{
		PieceWhiteKing:   0x0800000000000000,
		PieceWhiteQueen:  0x0000008000000000,
		PieceWhiteRook:   0x8000000000010000,
		PieceWhiteBishop: 0x2000100000000000,
		PieceWhiteKnight: 0x4000040000000000,
		PieceWhitePawn:   0x00D6002802000000,
		PieceBlackKing:   0x0000000000000004,
		PieceBlackQueen:  0x0000000000000010,
		PieceBlackRook:   0x0000000000000081,
		PieceBlackBishop: 0x0000000000000820,
		PieceBlackKnight: 0x0000000000240000,
		PieceBlackPawn:   0x000000004018A500,
	})

	testBoardStart = NewBoard(map[Piece]Bitboard{
		PieceWhiteKing:   0x0800000000000000,
		PieceWhiteQueen:  0x1000000000000000,
		PieceWhiteRook:   0x8100000000000000,
		PieceWhiteBishop: 0x2400000000000000,
		PieceWhiteKnight: 0x4200000000000000,
		PieceWhitePawn:   0x00FF000000000000,
		PieceBlackKing:   0x0000000000000008,
		PieceBlackQueen:  0x0000000000000010,
		PieceBlackRook:   0x0000000000000081,
		PieceBlackBishop: 0x0000000000000024,
		PieceBlackKnight: 0x0000000000000042,
		PieceBlackPawn:   0x000000000000FF00,
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
			fmt.Sprintf("required %d parts but got 7", len(ranks))},
		{
			"too many parts",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR/extra-part",
			nil,
			fmt.Sprintf("required %d parts but got 9", len(ranks)),
		},
		{
			"invalid piece type FEN",
			"rnbqkbnr/pppXpppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			"part #1, byte #3, NewPieceFromFEN(\"X\"): unknown FEN",
		},
		{
			"not enough files",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPP/RNBQKBNR",
			nil,
			fmt.Sprintf("required %d files but got 7 in part #6", len(files)),
		},
		{
			"too many files",
			"rrnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			fmt.Sprintf("required %d files but got 9 in part #0", len(files)),
		},
		{
			"not enough offsets",
			"rnbqkbnr/pppppppp/8/8/6/8/PPPPPPPP/RNBQKBNR",
			nil,
			fmt.Sprintf("required %d files but got 6 in part #4", len(files)),
		},
		{
			"too many offsets",
			"rnbqkbnr/pppppppp/9/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			fmt.Sprintf("required %d files but got 9 in part #2", len(files)),
		},
		{
			"not enough offsets and pieces",
			"r2bnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			fmt.Sprintf("required %d files but got 6 in part #0", len(files)),
		},
		{
			"too many offsets and pieces",
			"rnbqkbnr/p6ppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			fmt.Sprintf("required %d files but got 10 in part #1", len(files)),
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

func TestBoardGetColorBitboard(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		board          *Board
		colorBitboards map[Color]Bitboard
	}{

		{"start", testBoardStart, map[Color]Bitboard{ColorWhite: 0xFFFF000000000000, ColorBlack: 0x000000000000FFFF}},
		{"harder", testBoardHarder, map[Color]Bitboard{ColorWhite: 0xE8D614A802010000, ColorBlack: 0x00000000403CADB5}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for color, expectedBitboard := range test.colorBitboards {
				gotBitboard := test.board.GetColorBitboard(color)
				if gotBitboard != expectedBitboard {
					t.Fatalf("GetColorBitboard(%d) expected bitboard 0x%X but got 0x%X", color, expectedBitboard, gotBitboard)
				}
			}
		})
	}
}

func TestBoardGetOccupiedBitboard(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		board    *Board
		bitboard Bitboard
	}{
		{"start", testBoardStart, 0xFFFF00000000FFFF},
		{"harder", testBoardHarder, 0xE8D614A8423DADB5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			bitboard := test.board.GetOccupiedBitboard()
			if bitboard != test.bitboard {
				t.Fatalf("GetOccupiedBitboard() expected bitboard 0x%X but got 0x%X", test.bitboard, bitboard)
			}
		})
	}
}
