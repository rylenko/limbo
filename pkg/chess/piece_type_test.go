package chess

import "testing"

func TestNewPieceTypeFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		fen       string
		pieceType PieceType
		errString string
	}{
		{"k", PieceTypeBlackKing, ""},
		{"q", PieceTypeBlackQueen, ""},
		{"r", PieceTypeBlackRook, ""},
		{"b", PieceTypeBlackBishop, ""},
		{"n", PieceTypeBlackKnight, ""},
		{"p", PieceTypeBlackPawn, ""},
		{"K", PieceTypeWhiteKing, ""},
		{"Q", PieceTypeWhiteQueen, ""},
		{"R", PieceTypeWhiteRook, ""},
		{"B", PieceTypeWhiteBishop, ""},
		{"N", PieceTypeWhiteKnight, ""},
		{"P", PieceTypeWhitePawn, ""},
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

			pieceType, err := NewPieceTypeFromFEN(test.fen)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewPieceTypeFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if pieceType != test.pieceType {
				t.Fatalf("NewPieceTypeFromFEN(%q) expected %d but got %d", test.fen, test.pieceType, pieceType)
			}
		})
	}
}
