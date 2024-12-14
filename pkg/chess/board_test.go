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
			fmt.Sprintf("required %d parts separated by \"/\" but got 7", len(ranks))},
		{
			"too many parts",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR/extra-part",
			nil,
			fmt.Sprintf("required %d parts separated by \"/\" but got 9", len(ranks)),
		},
		{
			"invalid piece type FEN",
			"rnbqkbnr/pppXpppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			"part #1, byte #3, NewPieceFromFEN('X'): unknown FEN",
		},
		{"not enough files", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPP/RNBQKBNR", nil, "invalid files count in part #6"},
		{
			"too many files",
			"rrnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			"part #0, byte #8, NewSquare(Rank8, <unknown File=9>): NewSquaresOfFile(<unknown File=9>): unknown file",
		},
		{"not enough offsets", "rnbqkbnr/pppppppp/8/8/6/8/PPPPPPPP/RNBQKBNR", nil, "invalid files count in part #4"},
		{"too many offsets", "rnbqkbnr/pppppppp/9/8/8/8/PPPPPPPP/RNBQKBNR", nil, "invalid files count in part #2"},
		{"not enough offsets and pieces", "r2bnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", nil, "invalid files count in part #0"},
		{
			"too many offsets and pieces",
			"rnbqkbnr/p6ppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			"part #1, byte #3, NewSquare(Rank7, <unknown File=9>): NewSquaresOfFile(<unknown File=9>): unknown file",
		},
		{
			"zero offset",
			"rnbqkbnr/pppppppp/08/8/8/8/PPPPPPPP/RNBQKBNR",
			nil,
			"part #2, byte #0, NewPieceFromFEN('0'): unknown FEN",
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

	type testCase struct {
		color     Color
		bitboard  Bitboard
		errString string
	}

	tests := []struct {
		name  string
		board *Board
		cases []testCase
	}{

		{"start", testBoardStart, []testCase{
			{ColorWhite, 0xFFFF000000000000, ""},
			{ColorBlack, 0x000000000000FFFF, ""},
			{ColorNil, BitboardNil, "NewPiecesOfColor(ColorNil): no pieces"},
			{Color(123), BitboardNil, "NewPiecesOfColor(<unknown Color=123>): unknown color"},
		}},
		{"harder", testBoardHarder, []testCase{
			{ColorWhite, 0xE8D614A802010000, ""},
			{ColorBlack, 0x00000000403CADB5, ""},
			{ColorNil, BitboardNil, "NewPiecesOfColor(ColorNil): no pieces"},
			{Color(3), BitboardNil, "NewPiecesOfColor(<unknown Color=3>): unknown color"},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for _, casee := range test.cases {
				bitboard, err := test.board.GetColorBitboard(casee.color)
				if (err == nil && casee.errString != "") || (err != nil && err.Error() != casee.errString) {
					t.Fatalf("GetColorBitboard(%s) expected error %q but got %q", casee.color, casee.errString, err)
				}

				if bitboard != casee.bitboard {
					t.Fatalf("GetColorBitboard(%s) expected bitboard 0x%X but got 0x%X", casee.color, casee.bitboard, bitboard)
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

			bitboard, err := test.board.GetOccupiedBitboard()
			if err != nil {
				t.Fatalf("GetOccupiedBitboard() expected no error, but got %q", err)
			}

			if bitboard != test.bitboard {
				t.Fatalf("GetOccupiedBitboard() expected bitboard 0x%X but got 0x%X", test.bitboard, bitboard)
			}
		})
	}
}
