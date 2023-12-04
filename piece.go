package main

type Piece struct {
	pieceType  string
	pieceColor string
}

func (p *Piece) GetType() string {
	return p.pieceType + p.pieceColor
}
