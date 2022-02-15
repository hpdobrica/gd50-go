package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
	radius  float64
	x      float64
	y      float64
  dx     float64
  dy     float64
}

func NewBall(x, y,radius float64) Ball {
  return Ball {
    x: x,
    y:y,
    radius: radius,
  }
}

func (b Ball) IsColliding(p Paddle) bool {
  if (b.x < p.x + p.width && p.x < b.x + b.radius) && (b.y < p.y + p.height && p.y < b.y + b.radius) {
    return true
  }
  return false
}

func (b *Ball) Reset() {
  b.x = screenWidth/2- (b.radius/2)
  b.y = screenHeight/2 - (b.radius/2)
}

func (b *Ball) Update() {
    b.x += b.dx * dt
    b.y += b.dy * dt

}

func (b Ball) Draw(screen *ebiten.Image) {
  ebitenutil.DrawRect(screen, b.x, b.y, b.radius, b.radius, color.White)

}
