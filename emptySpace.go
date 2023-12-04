package main

type EmptySpace struct {
	Piece
	pieceColor string
}

func (es *EmptySpace) Init() {
	es.pieceType = "  "
}

func (es *EmptySpace) getType() string {
	return es.pieceType
}
