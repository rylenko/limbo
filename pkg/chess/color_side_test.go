package chess

import "testing"

func TestNewColorSideFromFEN(t *testing.T) {
	t.Parallel()

	validFENs := map[string]ColorSide{
		"k": ColorSideBlackKing,
		"q": ColorSideBlackQueen,
		"K": ColorSideWhiteKing,
		"Q": ColorSideWhiteQueen,
	}

	t.Run("valid FENs", func(t *testing.T) {
		t.Parallel()

		for fen, expectedPieceType := range validFENs {
			gotPieceType, err := NewColorSideFromFEN(fen)
			if err != nil {
				t.Fatalf("NewColorSideFromFEN(%q): %v", fen, err)
			}

			if gotPieceType != expectedPieceType {
				t.Fatalf("NewColorSideFromFEN(%q) expected %d but got %d", fen, expectedPieceType, gotPieceType)
			}
		}
	})

	const expectedErrString = "unknown FEN"
	invalidFENs := []string{"x", "y", "z", "xyz", "qwe", "kK"}

	t.Run("invalid FENs", func(t *testing.T) {
		t.Parallel()

		for _, fen := range invalidFENs {
			if _, err := NewColorSideFromFEN(fen); err == nil || err.Error() != expectedErrString {
				t.Fatalf("NewColorSideFromFEN(%q) expected error %q but got %q", fen, expectedErrString, err)
			}
		}
	})
}
