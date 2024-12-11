package chess

import "testing"

func TestNewColorSideFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		fen       string
		colorSide ColorSide
		errString string
	}{
		{"k", ColorSideBlackKing, ""},
		{"q", ColorSideBlackQueen, ""},
		{"K", ColorSideWhiteKing, ""},
		{"Q", ColorSideWhiteQueen, ""},
		{"x", ColorSideNil, "unknown FEN"},
		{"-", ColorSideNil, "unknown FEN"},
		{"kK", ColorSideNil, "unknown FEN"},
	}

	for _, test := range tests {
		t.Run(test.fen, func(t *testing.T) {
			t.Parallel()

			colorSide, err := NewColorSideFromFEN(test.fen)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewColorSideFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if colorSide != test.colorSide {
				t.Fatalf("NewColorSideFromFEN(%q) expected %q but got %q", test.fen, test.colorSide, colorSide)
			}
		})
	}
}

func TestColorSideString(t *testing.T) {
	t.Parallel()

	tests := []struct{
		colorSide ColorSide
		str string
	}{
		{ColorSideNil, "ColorSideNil"},
		{ColorSideBlackKing, "ColorSideBlackKing"},
		{ColorSideBlackQueen, "ColorSideBlackQueen"},
		{ColorSideWhiteKing, "ColorSideWhiteKing"},
		{ColorSideWhiteQueen, "ColorSideWhiteQueen"},
		{ColorSide(123), "<unknown ColorSide>"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			str := test.colorSide.String()
			if str != test.str {
				t.Fatalf("%s.String() expected %q but got %q", test.str, test.str, str)
			}
		})
	}
}
