package chess

import "testing"

func TestNewPieceTypeFromFEN(t *testing.T) {
	t.Parallel()

	validFENs := map[string]PieceType{
		"k": PieceTypeBlackKing,
		"q": PieceTypeBlackQueen,
		"r": PieceTypeBlackRook,
		"b": PieceTypeBlackBishop,
		"n": PieceTypeBlackKnight,
		"p": PieceTypeBlackPawn,
		"K": PieceTypeWhiteKing,
		"Q": PieceTypeWhiteQueen,
		"R": PieceTypeWhiteRook,
		"B": PieceTypeWhiteBishop,
		"N": PieceTypeWhiteKnight,
		"P": PieceTypeWhitePawn,
	}

	t.Run("valid FENs", func(t *testing.T) {
		t.Parallel()

		for fen, expectedPieceType := range validFENs {
			gotPieceType, err := NewPieceTypeFromFEN(fen)
			if err != nil {
				t.Fatalf("NewPieceTypeFromFEN(%q): %v", fen, err)
			}

			if gotPieceType != expectedPieceType {
				t.Fatalf("NewPieceTypeFromFEN(%q) expected %d but got %d", fen, expectedPieceType, gotPieceType)
			}
		}
	})

	const expectedErrString = "unknown FEN"
	invalidFENs := []string{"x", "y", "z", "xyz", "qwe", "kK"}

	t.Run("invalid FENs", func(t *testing.T) {
		t.Parallel()

		for _, fen := range invalidFENs {
			if _, err := NewPieceTypeFromFEN(fen); err == nil || err.Error() != expectedErrString {
				t.Fatalf("NewPieceTypeFromFEN(%q) expected error %q but got %q", fen, expectedErrString, err)
			}
		}
	})
}
