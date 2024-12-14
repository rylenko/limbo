package chess

import (
	"errors"
	"fmt"
)

// Color represents chess colors.
type Color uint8

const (
	ColorNil Color = iota
	ColorBlack
	ColorWhite
)

// NewColorFromFEN parses FEN to corresponding color or returns an error.
//
// FEN argument examples: "b", "w".
func NewColorFromFEN(fen string) (Color, error) {
	switch fen {
	case "b":
		return ColorBlack, nil
	case "w":
		return ColorWhite, nil
	default:
		return ColorNil, errors.New("unknown FEN")
	}
}

// Opposite returns opposite of current color.
func (color Color) Opposite() (Color, error) {
	switch color {
	case ColorBlack:
		return ColorWhite, nil
	case ColorWhite:
		return ColorBlack, nil
	case ColorNil:
		return ColorNil, errors.New("no opposite")
	default:
		return ColorNil, errors.New("unknown color")
	}
}

// String returns string representation of current color.
func (color Color) String() string {
	switch color {
	case ColorBlack:
		return "ColorBlack"
	case ColorWhite:
		return "ColorWhite"
	case ColorNil:
		return "ColorNil"
	default:
		return fmt.Sprintf("<unknown Color=%d>", color)
	}
}
