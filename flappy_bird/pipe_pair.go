package main

import "github.com/hajimehoshi/ebiten/v2"

type PipePair struct {
	top    *Pipe
	bottom *Pipe
	y      float64
	gap    int
}

func NewPipePair(y float64) PipePair {
	topPipe := NewPipe(y, "top")
	bottomPipe := NewPipe(y+90, "bottom")
	return PipePair{
		top:    &topPipe,
		bottom: &bottomPipe,
	}
}

func (p PipePair) Draw(screen *ebiten.Image) {
	p.top.Draw(screen)
	p.bottom.Draw(screen)
}

func (p *PipePair) Update() {
	p.top.Update()
	p.bottom.Update()

}
