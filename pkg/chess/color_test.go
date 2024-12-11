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
		{"B", ColorNil, "unknown FEN"},
		{"W", ColorNil, "unknown FEN"},
		{"-", ColorNil, "unknown FEN"},
		{"wb", ColorNil, "unknown FEN"},
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

func TestColorOpposite(t *testing.T) {
	t.Parallel()

	tests := []struct {
		color     Color
		opposite  Color
		errString string
	}{
		{ColorBlack, ColorWhite, ""},
		{ColorWhite, ColorBlack, ""},
		{ColorNil, ColorNil, "no opposite"},
		{Color(123), ColorNil, "unknown color"},
	}

	for _, test := range tests {
		t.Run(test.color.String(), func(t *testing.T) {
			t.Parallel()

			opposite, err := test.color.Opposite()
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("%s.Opposite() expected error %q but got %q", test.color, test.errString, err.Error())
			}

			if opposite != test.opposite {
				t.Fatalf("%s.Opposite() expected %s but got %s", test.color, test.opposite, opposite)
			}
		})
	}
}

func TestColorString(t *testing.T) {
	t.Parallel()

	tests := []struct{
		color Color
		str string
	}{
		{ColorNil, "ColorNil"},
		{ColorBlack, "ColorBlack"},
		{ColorWhite, "ColorWhite"},
		{Color(123), "<unknown Color>"},
	}

	for _, test := range tests {
		t.run(test.str, func(t *testing.T) {
			t.Parallel()

			str := test.color.String()
			if str != test.str {
				t.Fatalf("%s.String() expected %q but got %q", test.str, test.str, str)
			}
		})
	}
}
