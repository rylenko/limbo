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

// NewColorSideFromFEN parses FEN to corresponding ColorSide.
//
// FEN argument examples: "k", "q", "K", "Q".
func NewColorSideFromFEN(fen string) (ColorSide, error) {
	switch fen {
	case "k":
		return ColorSideBlackKing, nil
	case "q":
		return ColorSideBlackQueen, nil
	case "K":
		return ColorSideWhiteKing, nil
	case "Q":
		return ColorSideWhiteQueen, nil
	default:
		return 0, errors.New("unknown FEN")
	}
}
