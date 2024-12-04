package chess

import (
	"reflect"
	"testing"
)

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
			CastlingRights(map[ColorSide]struct{}{
				ColorSideBlackKing:  struct{}{},
				ColorSideBlackQueen: struct{}{},
				ColorSideWhiteKing:  struct{}{},
				ColorSideWhiteQueen: struct{}{},
			}),
			"",
		},
		{
			"black king side and white queen side",
			"Qk",
			CastlingRights(map[ColorSide]struct{}{
				ColorSideBlackKing:  struct{}{},
				ColorSideWhiteQueen: struct{}{},
			}),
			"",
		},
		{
			"no rights",
			"-",
			CastlingRights(map[ColorSide]struct{}{}),
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
			"castling right \"k\" already set",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rights, err := NewCastlingRightsFromFEN(test.fen)
			if (test.errString == "" && err != nil) || (test.errString != "" && (err == nil || test.errString != err.Error())) {
				t.Fatalf("NewCastlingRightsFromFEN(%q) expected error %q but got %q", test.fen, test.errString, err)
			}

			if !reflect.DeepEqual(rights, test.rights) {
				t.Fatalf("NewCastlingRightsFromFEN(%q) expected %+v but got %+v", test.fen, test.rights, rights)
			}
		})
	}
}
