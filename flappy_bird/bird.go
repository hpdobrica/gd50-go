package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const gravity = 20

type Bird struct {
	image  *ebiten.Image
	x      int
	y      int
	width  int
	height int

	dy float64
}

func NewBird() Bird {
	image := assets.getImage("assets/bird.png")
	width, height := image.Size()
	return Bird{
		image:  image,
		x:      screenWidth/2 - width/2,
		y:      screenHeight/2 - height/2,
		width:  width,
		height: height,
		dy:     0,
	}

}

func (b Bird) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.image, opts)

}

func (b *Bird) Update() {
	b.dy = b.dy + gravity*dt

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b.dy = -5
	}

	b.y = b.y + int(b.dy)
}
