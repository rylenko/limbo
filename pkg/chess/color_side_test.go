package chess

import "testing"

func TestNewColorSideFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		fen       string
		colorSide ColorSide
		errString string
	}{
		{"K", ColorSideWhiteKing, ""},
		{"Q", ColorSideWhiteQueen, ""},
		{"k", ColorSideBlackKing, ""},
		{"q", ColorSideBlackQueen, ""},
		{"x", ColorSide(0), "unknown FEN"},
		{"abc", ColorSide(0), "unknown FEN"},
		{"kK", ColorSide(0), "unknown FEN"},
	}

	for _, test := range tests {
		t.Run(test.fen, func(t *testing.T) {
			t.Parallel()

			colorSide, err := NewColorSideFromFEN(test.fen)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewColorSideFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if colorSide != test.colorSide {
				t.Fatalf("NewColorSideFromFEN(%q) expected %d but got %d", test.fen, test.colorSide, colorSide)
			}
		})
	}
}
