package chess

import "errors"

// ColorSide represents chess board's queen and king sides for each color.
type ColorSide uint8

const (
	ColorSideBlackKing ColorSide = iota
	ColorSideBlackQueen
	ColorSideWhiteKing
	ColorSideWhiteQueen
)

// Mapping of FENs to corresponding ColorSide.
var colorSideFENMap = map[string]ColorSide{
	"k": ColorSideBlackKing,
	"q": ColorSideBlackQueen,
	"K": ColorSideWhiteKing,
	"Q": ColorSideWhiteQueen,
}

// NewColorSideFromFEN parses FEN to corresponding ColorSide.
func NewColorSideFromFEN(fen string) (ColorSide, error) {
	side, ok := colorSideFENMap[fen]
	if !ok {
		return side, errors.New("unknown FEN")
	}
	return side, nil
}
