package main

type Piece struct {
	pieceType string
}

func (p *Piece) getType() string {
	return p.pieceType
}
