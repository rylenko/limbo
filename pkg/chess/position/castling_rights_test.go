package position

import (
	"reflect"
	"testing"
)

func TestNewCastlingRightsFromFEN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		fen       string
		rights    CastlingRights
		errString string
	}{
		{
			"kqKQ",
			CastlingRights([]ColorSide{ColorSideBlackKing, ColorSideBlackQueen, ColorSideWhiteKing, ColorSideWhiteQueen}),
			"",
		},
		{"qK", CastlingRights([]ColorSide{ColorSideBlackQueen, ColorSideWhiteKing}), ""},
		{"-", nil, ""},
		{"-k", nil, "NewColorSideFromFEN('-'): unknown FEN"},
		{"o", nil, "NewColorSideFromFEN('o'): unknown FEN"},
		{"kk", nil, "duplicate of 'k' found"},
	}

	for _, test := range tests {
		t.Run(test.fen, func(t *testing.T) {
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
