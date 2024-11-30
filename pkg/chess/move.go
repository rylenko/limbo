package chess

type Move struct {
	from  Square
	to    Square
	promo Role
}

func NewMove(from Square, to Square, promo Role) Move {
	return Move{
		from:  from,
		to:    to,
		promo: promo,
	}
}
