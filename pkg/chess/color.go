package chess

import "errors"

// Color represents chess colors.
type Color uint8

const (
	ColorWhite Color = iota
	ColorBlack
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
		return 0, errors.New("unknown FEN")
	}
}
