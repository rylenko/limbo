package chess

// Move represents single chess move.
type Move struct {
	from  Square
	to    Square
	promo Role
}

// NewMove creates a new Move with passed parameters.
func NewMove(from Square, to Square, promo Role) Move {
	return Move{
		from:  from,
		to:    to,
		promo: promo,
	}
}
