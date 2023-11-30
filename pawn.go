package main

//type piece_type string

type Pawn struct {
	Piece
}

func (p *Pawn) Init() {
	p.pieceType = "p"
}

func (p *Pawn) getPawn() string {
	return p.pieceType
}
