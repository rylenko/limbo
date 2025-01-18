package position

import (
	"errors"
	"fmt"
)

// ColorSide represents chess board's queen and king sides for each color.
type ColorSide uint8

const (
	ColorSideNil ColorSide = iota
	ColorSideBlackKing
	ColorSideBlackQueen
	ColorSideWhiteKing
	ColorSideWhiteQueen
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
		return ColorSideNil, errors.New("unknown FEN")
	}
}

// String returns string representation of current color side.
func (colorSide ColorSide) String() string {
	switch colorSide {
	case ColorSideWhiteKing:
		return "ColorSideWhiteKing"
	case ColorSideWhiteQueen:
		return "ColorSideWhiteQueen"
	case ColorSideBlackKing:
		return "ColorSideBlackKing"
	case ColorSideBlackQueen:
		return "ColorSideBlackQueen"
	case ColorSideNil:
		return "ColorSideNil"
	default:
		return fmt.Sprintf("<unknown ColorSide=%d>", colorSide)
	}
}
