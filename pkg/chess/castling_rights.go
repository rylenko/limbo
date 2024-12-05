package chess

import "fmt"

// CastlingRights is a set of color sides available for castling.
type CastlingRights map[ColorSide]struct{}

// NewCastlingRights creates a new set of castling rights.
func NewCastlingRights(colorSides ...ColorSide) CastlingRights {
	rights := make(map[ColorSide]struct{})

	for _, colorSide := range colorSides {
		rights[colorSide] = struct{}{}
	}

	return CastlingRights(rights)
}

// NewCastlingRightsFromFEN parses FEN to CastlingRights structure.
//
// FEN argument example: "kqKQ".
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
			return nil, fmt.Errorf("duplicate of %q found", fenByteString)
		}

		rights[colorSide] = struct{}{}
	}

	return CastlingRights(rights), nil
}

// Contains check presence of passed color side in castling rights.
func (rights CastlingRights) Contains(colorSide ColorSide) bool {
	_, ok := rights[colorSide]
	return ok
}
