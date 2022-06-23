package flyweight12

var units = map[int]*ChessPieceUnit{
	1: {
		1, "车", "red",
	},
	2: {
		2, "炮", "red",
	},
}

type ChessPieceUnit struct {
	ID    uint
	Name  string
	Color string
}

func NewChessPieceUnit(ID int) *ChessPieceUnit {
	return units[ID]
}

type ChessPiece struct {
	Unit *ChessPieceUnit
	X    int
	Y    int
}

type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

func NewChessBoard() *ChessBoard {
	board := &ChessBoard{chessPieces: map[int]*ChessPiece{}}
	for id := range units {
		board.chessPieces[id] = &ChessPiece{
			Unit: NewChessPieceUnit(id),
			X:    0,
			Y:    0,
		}
	}
	return board
}

func (c *ChessBoard) Move(id, x, y int) {
	c.chessPieces[id].X = x
	c.chessPieces[id].Y = y

}
