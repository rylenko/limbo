package chess

import "testing"

func TestNewColorFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		fen       string
		color     Color
		errString string
	}{
		{"b", ColorBlack, ""},
		{"w", ColorWhite, ""},
		{"B", 0, "unknown FEN"},
		{"W", 0, "unknown FEN"},
		{"-", 0, "unknown FEN"},
		{"wb", 0, "unknown FEN"},
	}

	for _, test := range tests {
		t.Run(test.fen, func(t *testing.T) {
			t.Parallel()

			color, err := NewColorFromFEN(test.fen)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewColorFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err.Error())
			}

			if color != test.color {
				t.Fatalf("NewColorFromFEN(%q) expected %d but got %d", test.fen, test.color, color)
			}
		})
	}
}
