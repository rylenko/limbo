package chess

import (
	"reflect"
	"testing"
)

func TestNewCastlingRights(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		colorSides        []ColorSide
		castlingRightsLen int
	}{
		{
			"all",
			[]ColorSide{ColorSideWhiteQueen, ColorSideWhiteKing, ColorSideBlackQueen, ColorSideBlackKing},
			4,
		},
		{
			"empty",
			[]ColorSide{},
			0,
		},
		{
			"duplicates",
			[]ColorSide{ColorSideWhiteQueen, ColorSideWhiteKing, ColorSideWhiteQueen, ColorSideWhiteQueen},
			2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rights := NewCastlingRights(test.colorSides...)
			if len(rights) != test.castlingRightsLen {
				t.Fatalf("NewCastlingRights(%v) len expected %d but got %d", test.colorSides, test.castlingRightsLen, len(rights))
			}

			for _, colorSide := range test.colorSides {
				if !rights.Contains(colorSide) {
					t.Fatalf("NewCastlingRights(%v) does not contains %d", test.colorSides, colorSide)
				}
			}
		})
	}
}

func TestNewCastlingRightsFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		fen       string
		rights    CastlingRights
		errString string
	}{
		{
			"all",
			"kqKQ",
			NewCastlingRights(ColorSideBlackKing, ColorSideBlackQueen, ColorSideWhiteKing, ColorSideWhiteQueen),
			"",
		},
		{
			"black king side and white queen side",
			"Qk",
			NewCastlingRights(ColorSideBlackKing, ColorSideWhiteQueen),
			"",
		},
		{
			"no rights",
			"-",
			NewCastlingRights(),
			"",
		},
		{
			"right with no rights",
			"-k",
			nil,
			"NewColorSideFromFEN(\"-\"): unknown FEN",
		},
		{
			"unknown color side FEN",
			"o",
			nil,
			"NewColorSideFromFEN(\"o\"): unknown FEN",
		},
		{
			"duplicate rights",
			"kk",
			nil,
			"duplicate of \"k\" found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rights, err := NewCastlingRightsFromFEN(test.fen)
			if (err == nil && test.errString != "") || (err != nil && err.Error() != test.errString) {
				t.Fatalf("NewCastlingRightsFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if !reflect.DeepEqual(rights, test.rights) {
				t.Fatalf("NewCastlingRightsFromFEN(%q) expected %+v but got %+v", test.fen, test.rights, rights)
			}
		})
	}
}
