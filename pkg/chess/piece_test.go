package chess

import (
	"slices"
	"testing"
)

func TestNewPiecesOfColor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		color  Color
		pieces []Piece
		errString string
	}{
		{"black", ColorBlack, []Piece{
			PieceBlackKing, PieceBlackQueen, PieceBlackRook, PieceBlackBishop, PieceBlackKnight, PieceBlackPawn}, ""},
		{"white", ColorWhite, []Piece{
			PieceWhiteKing, PieceWhiteQueen, PieceWhiteRook, PieceWhiteBishop, PieceWhiteKnight, PieceWhitePawn}, ""},
		{"nil", ColorNil, nil, "no pieces"},
		{"invalid", Color(123), nil, "unknown color"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			pieces, err := NewPiecesOfColor(test.color)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewPieceFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if !slices.Equal(pieces, test.pieces) {
				t.Fatalf("NewPiecesOfColor(%d) expected %v but got %v", test.color, test.pieces, pieces)
			}
		})
	}
}

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
		{"x", PieceNil, "unknown FEN"},
		{"xyz", PieceNil, "unknown FEN"},
		{"a", PieceNil, "unknown FEN"},
		{"abc", PieceNil, "unknown FEN"},
		{"", PieceNil, "unknown FEN"},
		{"-", PieceNil, "unknown FEN"},
	}

	for _, test := range tests {
		t.Run(test.fen, func(t *testing.T) {
			t.Parallel()

			piece, err := NewPieceFromFEN(test.fen)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewPieceFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if piece != test.piece {
				t.Fatalf("NewPieceFromFEN(%q) expected %s but got %s", test.fen, test.piece, piece)
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
		errString string
	}{
		{"whites", NewPiecesOfColor(ColorWhite), ColorWhite},
		{"blacks", NewPiecesOfColor(ColorBlack), ColorBlack},
		{"nil", []Piece{PieceNil}, ColorNil, "no color"},
		{"invalids", []Piece{Piece(123), Piece(100), Piece(200)}, ColorNil, "unknown piece"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for _, piece := range test.pieces {
				color, err := piece.Color();
				if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
					t.Fatalf("%s.Color() expected error %q but got %q", test.fen, test.errString, err)
				}

				if color != test.color {
					t.Fatalf("%s.Color() expected %s but got %s", piece, test.color, color)
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
		errString string
	}{
		{"kings", []Piece{PieceWhiteKing, PieceBlackKing}, RoleKing, ""},
		{"queens", []Piece{PieceWhiteQueen, PieceBlackQueen}, RoleQueen, ""},
		{"rooks", []Piece{PieceWhiteRook, PieceBlackRook}, RoleRook, ""},
		{"bishops", []Piece{PieceWhiteBishop, PieceBlackBishop}, RoleBishop, ""},
		{"knights", []Piece{PieceWhiteKnight, PieceBlackKnight}, RoleKnight, ""},
		{"pawns", []Piece{PieceWhitePawn, PieceBlackPawn}, RolePawn, ""},
		{"nil", []Piece{PieceNil}, RoleNil, "no role"},
		{"invalids", []Piece{Piece(123), Piece(111), Piece(222)}, RoleNil, "unknown piece"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			for _, piece := range test.pieces {
				role := piece.Role()
				if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
					t.Fatalf("%s.Role() expected error %q but got %q", piece, test.errString, err)
				}

				if role != test.role {
					t.Fatalf("%s.Role() expected %s but got %s", piece, test.role, role)
				}
			}
		})
	}
}

func TestPieceString(t *testing.T) {
	t.Parallel()

	tests := []struct{
		piece Piece
		str string
	} {
		{PieceNil, "PieceNil"},
		{PieceBlackKing, "PieceBlackKing"},
		{PieceBlackQueen, "PieceBlackQueen"},
		{PieceBlackRook, "PieceBlackRook"},
		{PieceBlackBishop, "PieceBlackBishop"},
		{PieceBlackKnight, "PieceBlackKnight"},
		{PieceBlackPawn, "PieceBlackPawn"},
		{PieceWhiteKing, "PieceWhiteKing"},
		{PieceWhiteQueen, "PieceWhiteQueen"},
		{PieceWhiteRook, "PieceWhiteRook"},
		{PieceWhiteBishop, "PieceWhiteBishop"},
		{PieceWhiteKnight, "PieceWhiteKnight"},
		{PieceWhitePawn, "PieceWhitePawn"},
		{Piece(123), "<unknown Piece>"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			t.Parallel()

			str := test.file.String()
			if str != test.str {
				t.Fatalf("%s.String() expected %q but got %q", test.str, test.str, str)
			}
		})
	}
}
