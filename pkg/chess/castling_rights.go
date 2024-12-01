package chess

import "fmt"

// CastlingRights is a set of color sides available for castling.
type CastlingRights map[ColorSide]struct{}

// NewCastlingRightsFromFEN parses FEN to CastlingRights structure.
func NewCastlingRightsFromFEN(fen string) (CastlingRights, error) {
	rights := make(map[ColorSide]struct{})

	if fen == "-" {
		return CastlingRights(rights), nil
	}

	for _, fenByte := range []byte(fen) {
		fenByteString := string(fenByte)

		colorSide, err := NewColorSideFromFEN(fenByteString)
		if err != nil {
			return nil, fmt.Errorf("NewColorSideFromFEN(%q): %w", fenByteString, err)
		}

		if _, ok := rights[colorSide]; ok {
			return nil, fmt.Errorf("castling right %q already set", fenByteString)
		}

		rights[colorSide] = struct{}{}
	}

	return CastlingRights(rights), nil
}
