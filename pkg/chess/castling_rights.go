package chess

import (
	"fmt"
	"slices"
)

// CastlingRights is a slice of color sides available for castling.
type CastlingRights []ColorSide

// NewCastlingRightsFromFEN parses FEN to CastlingRights structure.
//
// FEN argument example: "kqKQ".
func NewCastlingRightsFromFEN(fen string) (CastlingRights, error) {
	if fen == "-" {
		return CastlingRights(nil), nil
	}

	rights := make([]ColorSide, 0, len(fen))

	for _, fenByte := range []byte(fen) {
		fenByteString := string(fenByte)

		colorSide, err := NewColorSideFromFEN(fenByteString)
		if err != nil {
			return nil, fmt.Errorf("NewColorSideFromFEN(%q): %w", fenByteString, err)
		}

		if slices.Contains(rights, colorSide) {
			return nil, fmt.Errorf("duplicate of %q found", fenByteString)
		}

		rights = append(rights, colorSide)
	}

	return CastlingRights(rights), nil
}
