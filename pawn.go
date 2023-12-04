package main

type Pawn struct {
	Piece
}

func (p *Pawn) Init(pieceColor string) {
	p.pieceType = "p"
	p.pieceColor = pieceColor
}

func (p *Pawn) getPawn() string {
	return p.pieceType + p.pieceColor
}
