package move

var (

	// Contains bitboards of all possible antidiagonal destinations from passed origin.
	//
	// Note that the moves are raw, that is, for example, the piece moves can put their in checkmate. Moreover, these
	// moves do not exclude a collision with a piece of their own color.
	//
	//nolint:mnd // Magic numbers represents move Bitboards from specific square.
	moveAntidiagonalRawDestBitboards = map[Square]Bitboard{
		SquareA1: 0x0000000000000000,
		SquareB1: 0x0080000000000000,
		SquareC1: 0x0040800000000000,
		SquareD1: 0x0020408000000000,
		SquareE1: 0x0010204080000000,
		SquareF1: 0x0008102040800000,
		SquareG1: 0x0004081020408000,
		SquareH1: 0x0002040810204080,
		SquareA2: 0x4000000000000000,
		SquareB2: 0x2000800000000000,
		SquareC2: 0x1000408000000000,
		SquareD2: 0x0800204080000000,
		SquareE2: 0x0400102040800000,
		SquareF2: 0x0200081020408000,
		SquareG2: 0x0100040810204080,
		SquareH2: 0x0000020408102040,
		SquareA3: 0x2040000000000000,
		SquareB3: 0x1020008000000000,
		SquareC3: 0x0810004080000000,
		SquareD3: 0x0408002040800000,
		SquareE3: 0x0204001020408000,
		SquareF3: 0x0102000810204080,
		SquareG3: 0x0001000408102040,
		SquareH3: 0x0000000204081020,
		SquareA4: 0x1020400000000000,
		SquareB4: 0x0810200080000000,
		SquareC4: 0x0408100040800000,
		SquareD4: 0x0204080020408000,
		SquareE4: 0x0102040010204080,
		SquareF4: 0x0001020008102040,
		SquareG4: 0x0000010004081020,
		SquareH4: 0x0000000002040810,
		SquareA5: 0x0810204000000000,
		SquareB5: 0x0408102000800000,
		SquareC5: 0x0204081000408000,
		SquareD5: 0x0102040800204080,
		SquareE5: 0x0001020400102040,
		SquareF5: 0x0000010200081020,
		SquareG5: 0x0000000100040810,
		SquareH5: 0x0000000000020408,
		SquareA6: 0x0408102040000000,
		SquareB6: 0x0204081020008000,
		SquareC6: 0x0102040810004080,
		SquareD6: 0x0001020408002040,
		SquareE6: 0x0000010204001020,
		SquareF6: 0x0000000102000810,
		SquareG6: 0x0000000001000408,
		SquareH6: 0x0000000000000204,
		SquareA7: 0x0204081020400000,
		SquareB7: 0x0102040810200080,
		SquareC7: 0x0001020408100040,
		SquareD7: 0x0000010204080020,
		SquareE7: 0x0000000102040010,
		SquareF7: 0x0000000001020008,
		SquareG7: 0x0000000000010004,
		SquareH7: 0x0000000000000002,
		SquareA8: 0x0102040810204000,
		SquareB8: 0x0001020408102000,
		SquareC8: 0x0000010204081000,
		SquareD8: 0x0000000102040800,
		SquareE8: 0x0000000001020400,
		SquareF8: 0x0000000000010200,
		SquareG8: 0x0000000000000100,
		SquareH8: 0x0000000000000000,
	}

	// Contains bitboards of all possible diagonal destinations from passed origin.
	//
	// Note that the moves are raw, that is, for example, the piece moves can put him in checkmate. Moreover, these
	// moves do not exclude a collision with a piece of their own color.
	//
	//nolint:mnd // Magic numbers represents move Bitboards from specific square.
	moveDiagonalRawDestBitboards = map[Square]Bitboard{
		SquareA1: 0x0040201008040201,
		SquareB1: 0x0020100804020100,
		SquareC1: 0x0010080402010000,
		SquareD1: 0x0008040201000000,
		SquareE1: 0x0004020100000000,
		SquareF1: 0x0002010000000000,
		SquareG1: 0x0001000000000000,
		SquareH1: 0x0000000000000000,
		SquareA2: 0x0000402010080402,
		SquareB2: 0x8000201008040201,
		SquareC2: 0x4000100804020100,
		SquareD2: 0x2000080402010000,
		SquareE2: 0x1000040201000000,
		SquareF2: 0x0800020100000000,
		SquareG2: 0x0400010000000000,
		SquareH2: 0x0200000000000000,
		SquareA3: 0x0000004020100804,
		SquareB3: 0x0080002010080402,
		SquareC3: 0x8040001008040201,
		SquareD3: 0x4020000804020100,
		SquareE3: 0x2010000402010000,
		SquareF3: 0x1008000201000000,
		SquareG3: 0x0804000100000000,
		SquareH3: 0x0402000000000000,
		SquareA4: 0x0000000040201008,
		SquareB4: 0x0000800020100804,
		SquareC4: 0x0080400010080402,
		SquareD4: 0x8040200008040201,
		SquareE4: 0x4020100004020100,
		SquareF4: 0x2010080002010000,
		SquareG4: 0x1008040001000000,
		SquareH4: 0x0804020000000000,
		SquareA5: 0x0000000000402010,
		SquareB5: 0x0000008000201008,
		SquareC5: 0x0000804000100804,
		SquareD5: 0x0080402000080402,
		SquareE5: 0x8040201000040201,
		SquareF5: 0x4020100800020100,
		SquareG5: 0x2010080400010000,
		SquareH5: 0x1008040200000000,
		SquareA6: 0x0000000000004020,
		SquareB6: 0x0000000080002010,
		SquareC6: 0x0000008040001008,
		SquareD6: 0x0000804020000804,
		SquareE6: 0x0080402010000402,
		SquareF6: 0x8040201008000201,
		SquareG6: 0x4020100804000100,
		SquareH6: 0x2010080402000000,
		SquareA7: 0x0000000000000040,
		SquareB7: 0x0000000000800020,
		SquareC7: 0x0000000080400010,
		SquareD7: 0x0000008040200008,
		SquareE7: 0x0000804020100004,
		SquareF7: 0x0080402010080002,
		SquareG7: 0x8040201008040001,
		SquareH7: 0x4020100804020000,
		SquareA8: 0x0000000000000000,
		SquareB8: 0x0000000000008000,
		SquareC8: 0x0000000000804000,
		SquareD8: 0x0000000080402000,
		SquareE8: 0x0000008040201000,
		SquareF8: 0x0000804020100800,
		SquareG8: 0x0080402010080400,
		SquareH8: 0x8040201008040200,
	}

	// Contains bitboards of all possible horizontal destinations from passed origin.
	//
	// Note that the moves are raw, that is, for example, the piece moves can put him in checkmate. Moreover, these
	// moves do not exclude a collision with a piece of their own color.
	//
	//nolint:mnd // Magic numbers represents move Bitboards from specific square.
	moveHorizontalRawDestBitboards = map[Square]Bitboard{
		SquareA1: 0x7f00000000000000,
		SquareB1: 0xbf00000000000000,
		SquareC1: 0xdf00000000000000,
		SquareD1: 0xef00000000000000,
		SquareE1: 0xf700000000000000,
		SquareF1: 0xfb00000000000000,
		SquareG1: 0xfd00000000000000,
		SquareH1: 0xfe00000000000000,
		SquareA2: 0x007f000000000000,
		SquareB2: 0x00bf000000000000,
		SquareC2: 0x00df000000000000,
		SquareD2: 0x00ef000000000000,
		SquareE2: 0x00f7000000000000,
		SquareF2: 0x00fb000000000000,
		SquareG2: 0x00fd000000000000,
		SquareH2: 0x00fe000000000000,
		SquareA3: 0x00007f0000000000,
		SquareB3: 0x0000bf0000000000,
		SquareC3: 0x0000df0000000000,
		SquareD3: 0x0000ef0000000000,
		SquareE3: 0x0000f70000000000,
		SquareF3: 0x0000fb0000000000,
		SquareG3: 0x0000fd0000000000,
		SquareH3: 0x0000fe0000000000,
		SquareA4: 0x0000007f00000000,
		SquareB4: 0x000000bf00000000,
		SquareC4: 0x000000df00000000,
		SquareD4: 0x000000ef00000000,
		SquareE4: 0x000000f700000000,
		SquareF4: 0x000000fb00000000,
		SquareG4: 0x000000fd00000000,
		SquareH4: 0x000000fe00000000,
		SquareA5: 0x000000007f000000,
		SquareB5: 0x00000000bf000000,
		SquareC5: 0x00000000df000000,
		SquareD5: 0x00000000ef000000,
		SquareE5: 0x00000000f7000000,
		SquareF5: 0x00000000fb000000,
		SquareG5: 0x00000000fd000000,
		SquareH5: 0x00000000fe000000,
		SquareA6: 0x00000000007f0000,
		SquareB6: 0x0000000000bf0000,
		SquareC6: 0x0000000000df0000,
		SquareD6: 0x0000000000ef0000,
		SquareE6: 0x0000000000f70000,
		SquareF6: 0x0000000000fb0000,
		SquareG6: 0x0000000000fd0000,
		SquareH6: 0x0000000000fe0000,
		SquareA7: 0x0000000000007f00,
		SquareB7: 0x000000000000bf00,
		SquareC7: 0x000000000000df00,
		SquareD7: 0x000000000000ef00,
		SquareE7: 0x000000000000f700,
		SquareF7: 0x000000000000fb00,
		SquareG7: 0x000000000000fd00,
		SquareH7: 0x000000000000fe00,
		SquareA8: 0x000000000000007f,
		SquareB8: 0x00000000000000bf,
		SquareC8: 0x00000000000000df,
		SquareD8: 0x00000000000000ef,
		SquareE8: 0x00000000000000f7,
		SquareF8: 0x00000000000000fb,
		SquareG8: 0x00000000000000fd,
		SquareH8: 0x00000000000000fe,
	}

	// Contains bitboards of all possible king destinations from passed origin.
	//
	// Note that the moves are raw, that is, for example, the king moves can put him in checkmate. Moreover, these
	// moves do not exclude a collision with a piece of their own color.
	//
	//nolint:mnd // Magic numbers represents move Bitboards from specific square.
	moveKingRawDestBitboards = map[Square]Bitboard{
		SquareA1: 0x40c0000000000000,
		SquareB1: 0xa0e0000000000000,
		SquareC1: 0x5070000000000000,
		SquareD1: 0x2838000000000000,
		SquareE1: 0x141c000000000000,
		SquareF1: 0x0a0e000000000000,
		SquareG1: 0x0507000000000000,
		SquareH1: 0x0203000000000000,
		SquareA2: 0xc040c00000000000,
		SquareB2: 0xe0a0e00000000000,
		SquareC2: 0x7050700000000000,
		SquareD2: 0x3828380000000000,
		SquareE2: 0x1c141c0000000000,
		SquareF2: 0x0e0a0e0000000000,
		SquareG2: 0x0705070000000000,
		SquareH2: 0x0302030000000000,
		SquareA3: 0x00c040c000000000,
		SquareB3: 0x00e0a0e000000000,
		SquareC3: 0x0070507000000000,
		SquareD3: 0x0038283800000000,
		SquareE3: 0x001c141c00000000,
		SquareF3: 0x000e0a0e00000000,
		SquareG3: 0x0007050700000000,
		SquareH3: 0x0003020300000000,
		SquareA4: 0x0000c040c0000000,
		SquareB4: 0x0000e0a0e0000000,
		SquareC4: 0x0000705070000000,
		SquareD4: 0x0000382838000000,
		SquareE4: 0x00001c141c000000,
		SquareF4: 0x00000e0a0e000000,
		SquareG4: 0x0000070507000000,
		SquareH4: 0x0000030203000000,
		SquareA5: 0x000000c040c00000,
		SquareB5: 0x000000e0a0e00000,
		SquareC5: 0x0000007050700000,
		SquareD5: 0x0000003828380000,
		SquareE5: 0x0000001c141c0000,
		SquareF5: 0x0000000e0a0e0000,
		SquareG5: 0x0000000705070000,
		SquareH5: 0x0000000302030000,
		SquareA6: 0x00000000c040c000,
		SquareB6: 0x00000000e0a0e000,
		SquareC6: 0x0000000070507000,
		SquareD6: 0x0000000038283800,
		SquareE6: 0x000000001c141c00,
		SquareF6: 0x000000000e0a0e00,
		SquareG6: 0x0000000007050700,
		SquareH6: 0x0000000003020300,
		SquareA7: 0x0000000000c040c0,
		SquareB7: 0x0000000000e0a0e0,
		SquareC7: 0x0000000000705070,
		SquareD7: 0x0000000000382838,
		SquareE7: 0x00000000001c141c,
		SquareF7: 0x00000000000e0a0e,
		SquareG7: 0x0000000000070507,
		SquareH7: 0x0000000000030203,
		SquareA8: 0x000000000000c040,
		SquareB8: 0x000000000000e0a0,
		SquareC8: 0x0000000000007050,
		SquareD8: 0x0000000000003828,
		SquareE8: 0x0000000000001c14,
		SquareF8: 0x0000000000000e0a,
		SquareG8: 0x0000000000000705,
		SquareH8: 0x0000000000000302,
	}

	// Contains bitboards of all possible knight destinations from passed origin.
	//
	// Note that the moves are raw, that is, for example, the knight moves can put his king in checkmate. Moreover, these
	// moves do not exclude a collision with a piece of their own color.
	//
	//nolint:mnd // Magic numbers represents move Bitboards from specific square.
	moveKnightRawDestBitboards = map[Square]Bitboard{
		SquareA1: 0x0020400000000000,
		SquareB1: 0x0010a00000000000,
		SquareC1: 0x0088500000000000,
		SquareD1: 0x0044280000000000,
		SquareE1: 0x0022140000000000,
		SquareF1: 0x00110a0000000000,
		SquareG1: 0x0008050000000000,
		SquareH1: 0x0004020000000000,
		SquareA2: 0x2000204000000000,
		SquareB2: 0x100010a000000000,
		SquareC2: 0x8800885000000000,
		SquareD2: 0x4400442800000000,
		SquareE2: 0x2200221400000000,
		SquareF2: 0x1100110a00000000,
		SquareG2: 0x0800080500000000,
		SquareH2: 0x0400040200000000,
		SquareA3: 0x4020002040000000,
		SquareB3: 0xa0100010a0000000,
		SquareC3: 0x5088008850000000,
		SquareD3: 0x2844004428000000,
		SquareE3: 0x1422002214000000,
		SquareF3: 0x0a1100110a000000,
		SquareG3: 0x0508000805000000,
		SquareH3: 0x0204000402000000,
		SquareA4: 0x0040200020400000,
		SquareB4: 0x00a0100010a00000,
		SquareC4: 0x0050880088500000,
		SquareD4: 0x0028440044280000,
		SquareE4: 0x0014220022140000,
		SquareF4: 0x000a1100110a0000,
		SquareG4: 0x0005080008050000,
		SquareH4: 0x0002040004020000,
		SquareA5: 0x0000402000204000,
		SquareB5: 0x0000a0100010a000,
		SquareC5: 0x0000508800885000,
		SquareD5: 0x0000284400442800,
		SquareE5: 0x0000142200221400,
		SquareF5: 0x00000a1100110a00,
		SquareG5: 0x0000050800080500,
		SquareH5: 0x0000020400040200,
		SquareA6: 0x0000004020002040,
		SquareB6: 0x000000a0100010a0,
		SquareC6: 0x0000005088008850,
		SquareD6: 0x0000002844004428,
		SquareE6: 0x0000001422002214,
		SquareF6: 0x0000000a1100110a,
		SquareG6: 0x0000000508000805,
		SquareH6: 0x0000000204000402,
		SquareA7: 0x0000000040200020,
		SquareB7: 0x00000000a0100010,
		SquareC7: 0x0000000050880088,
		SquareD7: 0x0000000028440044,
		SquareE7: 0x0000000014220022,
		SquareF7: 0x000000000a110011,
		SquareG7: 0x0000000005080008,
		SquareH7: 0x0000000002040004,
		SquareA8: 0x0000000000402000,
		SquareB8: 0x0000000000a01000,
		SquareC8: 0x0000000000508800,
		SquareD8: 0x0000000000284400,
		SquareE8: 0x0000000000142200,
		SquareF8: 0x00000000000a1100,
		SquareG8: 0x0000000000050800,
		SquareH8: 0x0000000000020400,
	}

	// Contains bitboards of all possible vertical destinations from passed origin.
	//
	// Note that the moves are raw, that is, for example, the piece moves can put him in checkmate. Moreover, these
	// moves do not exclude a collision with a piece of their own color.
	//
	//nolint:mnd // Magic numbers represents move Bitboards from specific square.
	moveVerticalRawDestBitboards = map[Square]Bitboard{
		SquareA1: 0x0080808080808080,
		SquareB1: 0x0040404040404040,
		SquareC1: 0x0020202020202020,
		SquareD1: 0x0010101010101010,
		SquareE1: 0x0008080808080808,
		SquareF1: 0x0004040404040404,
		SquareG1: 0x0002020202020202,
		SquareH1: 0x0001010101010101,
		SquareA2: 0x8000808080808080,
		SquareB2: 0x4000404040404040,
		SquareC2: 0x2000202020202020,
		SquareD2: 0x1000101010101010,
		SquareE2: 0x0800080808080808,
		SquareF2: 0x0400040404040404,
		SquareG2: 0x0200020202020202,
		SquareH2: 0x0100010101010101,
		SquareA3: 0x8080008080808080,
		SquareB3: 0x4040004040404040,
		SquareC3: 0x2020002020202020,
		SquareD3: 0x1010001010101010,
		SquareE3: 0x0808000808080808,
		SquareF3: 0x0404000404040404,
		SquareG3: 0x0202000202020202,
		SquareH3: 0x0101000101010101,
		SquareA4: 0x8080800080808080,
		SquareB4: 0x4040400040404040,
		SquareC4: 0x2020200020202020,
		SquareD4: 0x1010100010101010,
		SquareE4: 0x0808080008080808,
		SquareF4: 0x0404040004040404,
		SquareG4: 0x0202020002020202,
		SquareH4: 0x0101010001010101,
		SquareA5: 0x8080808000808080,
		SquareB5: 0x4040404000404040,
		SquareC5: 0x2020202000202020,
		SquareD5: 0x1010101000101010,
		SquareE5: 0x0808080800080808,
		SquareF5: 0x0404040400040404,
		SquareG5: 0x0202020200020202,
		SquareH5: 0x0101010100010101,
		SquareA6: 0x8080808080008080,
		SquareB6: 0x4040404040004040,
		SquareC6: 0x2020202020002020,
		SquareD6: 0x1010101010001010,
		SquareE6: 0x0808080808000808,
		SquareF6: 0x0404040404000404,
		SquareG6: 0x0202020202000202,
		SquareH6: 0x0101010101000101,
		SquareA7: 0x8080808080800080,
		SquareB7: 0x4040404040400040,
		SquareC7: 0x2020202020200020,
		SquareD7: 0x1010101010100010,
		SquareE7: 0x0808080808080008,
		SquareF7: 0x0404040404040004,
		SquareG7: 0x0202020202020002,
		SquareH7: 0x0101010101010001,
		SquareA8: 0x8080808080808000,
		SquareB8: 0x4040404040404000,
		SquareC8: 0x2020202020202000,
		SquareD8: 0x1010101010101000,
		SquareE8: 0x0808080808080800,
		SquareF8: 0x0404040404040400,
		SquareG8: 0x0202020202020200,
		SquareH8: 0x0101010101010100,
	}
)

// Move represents single chess move.
type Move struct {
	origin    Square
	dest      Square
	tags      MoveTags
	promoRole Role
}

// NewMove creates a new Move with passed parameters.
func NewMove(origin, dest Square, tags MoveTags, promoRole Role) Move {
	return Move{
		origin:    origin,
		dest:      dest,
		tags:      tags,
		promoRole: promoRole,
	}
}

// NewMovesPromo creates new equal moves but with all different promotions.
//
// TODO: test.
func NewMovesPromo(origin, dest Square, tags MoveTags) []Move {
	moves := make([]Move, 0, len(rolePromos))

	for _, promoRole := range rolePromos {
		moves = append(moves, NewMove(origin, dest, tags, promoRole))
	}

	return moves
}

// MoveTag represents cached useful notes about move.
//
// TODO move MoveTag and MoveTags to separate files.
type MoveTag uint8

const (
	MoveTagCapture MoveTag = 1 << iota
	MoveTagCheck
	MoveTagEnPassantCapture
	MoveTagKingSideCastle
	MoveTagQueenSideCastle
)

// Move tags contains several MoveTag. The list is an unsigned integer. Each tag is assumed to occupy a separate bit.
//
// Zero value is ready to use.
//
// TODO move MoveTag and MoveTags to separate files.
type MoveTags uint8

const MoveTagsNil MoveTags = iota

// Set sets passed tag to the tags list.
func (tags *MoveTags) Set(tag MoveTag) {
	*tags |= MoveTags(tag)
}

// Contains checks that tags list contains passed tag.
func (tags MoveTags) Contains(tag MoveTag) bool {
	return tags&MoveTags(tag) > 0
}
