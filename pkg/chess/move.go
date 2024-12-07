package chess

// Move represents single chess move.
type Move struct {
	origin Square
	dest   Square
	promo  Role
}

// NewMove creates a new Move with passed parameters.
func NewMove(origin Square, dest Square, promo Role) Move {
	return Move{
		origin: origin,
		dest:   dest,
		promo:  promo,
	}
}
