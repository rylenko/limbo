package position

import (
	"fmt"
	"slices"
)

// CastlingRights is a slice of color sides available for castling.
type CastlingRights []ColorSide

// NewRightsFromFEN parses FEN to Rights structure.
//
// FEN argument example: "kqKQ".
func NewRightsFromFEN(fen string) (Rights, error) {
	if fen == "-" {
		return Rights(nil), nil
	}

	rights := make([]ColorSide, 0, len(fen))

	for _, bytee := range []byte(fen) {
		colorSide, err := NewColorSideFromFEN(string(bytee))
		if err != nil {
			return nil, fmt.Errorf("NewColorSideFromFEN(%q): %w", bytee, err)
		}

		if slices.Contains(rights, colorSide) {
			return nil, fmt.Errorf("duplicate of %q found", bytee)
		}

		rights = append(rights, colorSide)
	}

	return Rights(rights), nil
}
