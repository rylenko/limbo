package chess

import (
	"fmt"
	"reflect"
	"testing"
)

var testPositionStart = NewPosition(
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
	CastlingRights([]ColorSide{ColorSideWhiteKing, ColorSideWhiteQueen, ColorSideBlackKing, ColorSideBlackQueen}),
	nil,
	0,
	1,
)

func TestNewPositionStart(t *testing.T) {
	t.Parallel()

	position, err := NewPositionStart()
	if err != nil {
		t.Fatalf("NewPositionStart(): %v", err)
	}

	if !reflect.DeepEqual(position, testPositionStart) {
		t.Fatalf("NewPositionStart() expected %+v but got %+v", testPositionStart, position)
	}
}

func TestNewPositionFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		fen       string
		position  *Position
		errString string
	}{
		{"start", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", testPositionStart, ""},
		{
			"no full move number",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0",
			nil,
			fmt.Sprintf("FEN parts required %d but got 5", positionFENPartsCount),
		},
		{
			"invalid board",
			"rnbqkbnr/ppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			nil,
			fmt.Sprintf(
				"NewBoardFromFEN(\"rnbqkbnr/ppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR\"): required %d files but got 7 in part #1",
				filesCount),
		},
		{
			"invalid active color",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR x KQkq - 0 1",
			nil,
			"NewColorFromFEN(\"x\"): unknown FEN",
		},
		{
			"invalid castling rights",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQQq - 0 1",
			nil,
			"NewCastlingRightsFromFEN(\"KQQq\"): duplicate of \"Q\" found",
		},
		{
			"invalid En Passant",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq f2 0 1",
			nil,
			"NewSquareEnPassantFromFEN(\"f2\"): invalid rank 1",
		},
		{
			"invalid half move clock",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 300 1",
			nil,
			"half move clock is not uint8: strconv.ParseUint: parsing \"300\": value out of range",
		},
		{
			"invalid full move number",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 -1",
			nil,
			"full move number is not uint16: strconv.ParseUint: parsing \"-1\": invalid syntax",
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
