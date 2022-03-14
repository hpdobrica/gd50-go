package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const pipeScroll float64 = -60

const pipeHeightMin float64 = screenHeight / 4
const pipeHeightMax float64 = screenHeight - 50

type Pipe struct {
	image  *ebiten.Image
	x      float64
	y      float64
	width  int
	height int

	orientation string

	dy float64
}

func NewPipe(y float64, orientation string) Pipe {
	image := assets.getImage("assets/pipe.png")
	width, height := image.Size()
	return Pipe{
		image:       image,
		x:           screenWidth,
		y:           y,
		width:       width,
		height:      height,
		dy:          0,
		orientation: orientation,
	}

}

func (p Pipe) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	if p.orientation == "top" {
		opts.GeoM.Scale(1, -1)
		// opts.GeoM.Translate(0, float64(p.height))
	}
	opts.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(p.image, opts)

}

func (p *Pipe) Update() {

	p.x = p.x + pipeScroll*dt

}
