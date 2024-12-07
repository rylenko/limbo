package chess

import "errors"

// ColorSide represents chess board's queen and king sides for each color.
type ColorSide uint8

const (
	ColorSideWhiteKing ColorSide = iota
	ColorSideWhiteQueen
	ColorSideBlackKing
	ColorSideBlackQueen
)

var (
	colorSides = []ColorSide{ColorSideWhiteKing, ColorSideWhiteQueen, ColorSideBlackKing, ColorSideBlackQueen}

	// Mapping of FENs to corresponding ColorSide.
	colorSideFENMap = map[string]ColorSide{
		"k": ColorSideBlackKing,
		"q": ColorSideBlackQueen,
		"K": ColorSideWhiteKing,
		"Q": ColorSideWhiteQueen,
	}
)

// NewColorSideFromFEN parses FEN to corresponding ColorSide.
//
// FEN argument examples: "k", "q", "K", "Q".
func NewColorSideFromFEN(fen string) (ColorSide, error) {
	side, ok := colorSideFENMap[fen]
	if !ok {
		return side, errors.New("unknown FEN")
	}
	return side, nil
}
