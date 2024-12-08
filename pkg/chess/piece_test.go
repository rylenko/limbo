package chess

import (
	"slices"
	"testing"
)

func TestNewPieceFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		fen       string
		piece     Piece
		errString string
	}{
		{"k", PieceBlackKing, ""},
		{"q", PieceBlackQueen, ""},
		{"r", PieceBlackRook, ""},
		{"b", PieceBlackBishop, ""},
		{"n", PieceBlackKnight, ""},
		{"p", PieceBlackPawn, ""},
		{"K", PieceWhiteKing, ""},
		{"Q", PieceWhiteQueen, ""},
		{"R", PieceWhiteRook, ""},
		{"B", PieceWhiteBishop, ""},
		{"N", PieceWhiteKnight, ""},
		{"P", PieceWhitePawn, ""},
		{"x", 0, "unknown FEN"},
		{"xyz", 0, "unknown FEN"},
		{"a", 0, "unknown FEN"},
		{"abc", 0, "unknown FEN"},
		{"", 0, "unknown FEN"},
		{"-", 0, "unknown FEN"},
	}

	for _, test := range tests {
		t.Run(test.fen, func(t *testing.T) {
			t.Parallel()

			piece, err := NewPieceFromFEN(test.fen)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewPieceFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if piece != test.piece {
				t.Fatalf("NewPieceFromFEN(%q) expected %d but got %d", test.fen, test.piece, piece)
			}
		})
	}
}

func TestNewPiecesOfColor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		color  Color
		pieces []Piece
	}{
		{"black", ColorBlack, []Piece{
			PieceBlackKing, PieceBlackQueen, PieceBlackRook, PieceBlackBishop, PieceBlackKnight, PieceBlackPawn}},
		{"white", ColorWhite, []Piece{
			PieceWhiteKing, PieceWhiteQueen, PieceWhiteRook, PieceWhiteBishop, PieceWhiteKnight, PieceWhitePawn}},
		{"invalid", Color(123), nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			pieces := NewPiecesOfColor(test.color)
			if !slices.Equal(pieces, test.pieces) {
				t.Fatalf("NewPiecesOfColor(%d) expected %v but got %v", test.color, test.pieces, pieces)
			}
		})
	}
}

func TestPieceColor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		pieces []Piece
		color  Color
	}{
		{"whites", NewPiecesOfColor(ColorWhite), ColorWhite},
		{"blacks", NewPiecesOfColor(ColorBlack), ColorBlack},
		{"invalids", []Piece{Piece(123), Piece(100), Piece(200)}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for _, piece := range test.pieces {
				if color := piece.Color(); color != test.color {
					t.Fatalf("Piece(%d).Color() expected %d but got %d", piece, test.color, color)
				}
			}
		})
	}
}

func TestPieceRole(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		pieces []Piece
		role   Role
	}{
		{"kings", []Piece{PieceWhiteKing, PieceBlackKing}, RoleKing},
		{"queens", []Piece{PieceWhiteQueen, PieceBlackQueen}, RoleQueen},
		{"rooks", []Piece{PieceWhiteRook, PieceBlackRook}, RoleRook},
		{"bishops", []Piece{PieceWhiteBishop, PieceBlackBishop}, RoleBishop},
		{"knights", []Piece{PieceWhiteKnight, PieceBlackKnight}, RoleKnight},
		{"pawns", []Piece{PieceWhitePawn, PieceBlackPawn}, RolePawn},
		{"invalids", []Piece{Piece(123), Piece(111), Piece(222)}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for _, piece := range test.pieces {
				if role := piece.Role(); role != test.role {
					t.Fatalf("Piece(%d).Role() expected %d but got %d", piece, test.role, role)
				}
			}
		})
	}
}
