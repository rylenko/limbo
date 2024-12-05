package chess

import "errors"

// Color represents chess colors.
type Color int8

const (
	ColorWhite Color = iota
	ColorBlack
)

// Mapping of FENs to corresponding colors.
var colorFENMap = map[string]Color{
	"b": ColorBlack,
	"w": ColorWhite,
}

// NewColorFromFEN parses FEN to corresponding color or returns an error.
//
// FEN argument examples: "b", "w".
func NewColorFromFEN(fen string) (Color, error) {
	color, ok := colorFENMap[fen]
	if !ok {
		return color, errors.New("unknown FEN")
	}
	return color, nil
}
