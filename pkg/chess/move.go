package chess

// Move represents single chess move.
type Move struct {
	origin  Square
	dest    Square
	isPromo bool
}

// NewMove creates a new Move with passed parameters.
func NewMove(origin Square, dest Square, isPromo bool) Move {
	return Move{
		origin:  origin,
		dest:    dest,
		isPromo: isPromo,
	}
}
