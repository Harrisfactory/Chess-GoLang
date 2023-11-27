package main

type EmptySpace struct {
	Piece
}

func (es *EmptySpace) Init() {
	es.pieceType = " "
}

func (es *EmptySpace) getType() string {
	return es.pieceType
}
