package main

import (
	"math/rand"

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

	dy float64
}

func NewPipe() Pipe {
	image := assets.getImage("assets/pipe.png")
	width, height := image.Size()
	return Pipe{
		image:  image,
		x:      screenWidth,
		y:      float64(rand.Intn(int(pipeHeightMax-pipeHeightMin)) + int(pipeHeightMin)),
		width:  width,
		height: height,
		dy:     0,
	}

}

func (p Pipe) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(p.image, opts)

}

func (p *Pipe) Update() {

	p.x = p.x + pipeScroll*dt

}
