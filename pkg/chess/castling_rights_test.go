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
			CastlingRights([]ColorSide{ColorSideBlackKing, ColorSideBlackQueen, ColorSideWhiteKing, ColorSideWhiteQueen}),
			"",
		},
		{
			"black king side and white queen side",
			"qK",
			CastlingRights([]ColorSide{ColorSideBlackQueen, ColorSideWhiteKing}),
			"",
		},
		{
			"no rights",
			"-",
			CastlingRights([]ColorSide{}),
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
