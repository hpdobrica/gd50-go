package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const gravity = 20

type Bird struct {
	image  *ebiten.Image
	x      float64
	y      float64
	width  int
	height int

	dy float64
}

func NewBird() Bird {
	image := assets.getImage("assets/bird.png")
	width, height := image.Size()
	return Bird{
		image:  image,
		x:      float64(screenWidth/2 - width/2),
		y:      float64(screenHeight/2 - height/2),
		width:  width,
		height: height,
		dy:     0,
	}

}

func (b Bird) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(b.x, b.y)
	screen.DrawImage(b.image, opts)

}

func (b *Bird) Update() {
	b.dy = b.dy + gravity*dt

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b.dy = -5
	}

	b.y = b.y + b.dy
}

func (b Bird) Collides(pipe Pipe) bool {
	var generosity float64 = 2

	collides := aabbCollides(b.x+generosity, b.y+generosity, float64(b.width)-generosity*2, float64(b.height)-generosity*2, pipe.x, pipe.y, float64(pipe.width), float64(pipe.height))

	fmt.Println(collides, b.x, b.y, float64(b.width), float64(b.height), pipe.x, pipe.y, float64(pipe.width), float64(pipe.height))
	// true b.x:237 b.y:78.99999999999999 b.width38 b.heigth24 pipe.x:274 pipe.y:40 pipe.width70 pipe.heigth:288

	return collides
}
