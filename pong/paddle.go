package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Paddle struct {
	width  float64
	height float64
	x      float64
	y      float64
  dy     float64
}

func NewPaddle(x, y, width, height float64) Paddle {
  return Paddle{
    x: x,
    y:y,
    width: width,
    height: height,
  }
}

func (p *Paddle) Update() {
  if p.dy < 0 {
    p.y = math.Max(0, p.y + p.dy * dt)
  } else {
    p.y = math.Min(screenHeight - p.height, p.y + p.dy * dt)
  }
}

func (p Paddle) Draw(screen *ebiten.Image) {
  ebitenutil.DrawRect(screen, p.x, p.y, p.width, p.height, color.White)

}
