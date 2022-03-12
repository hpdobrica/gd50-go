package main

import (
	"bytes"
	"embed"
	"image"
	"log"

  _ "image/png"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var _assets embed.FS

type GameAssets struct{
  images map[string]*ebiten.Image
}

func NewGameAssets() GameAssets {
  return GameAssets{
    images: make(map[string]*ebiten.Image),
  }
}


func (ga *GameAssets) getImage(path string) *ebiten.Image {
  if image, ok := ga.images[path]; ok {
    return image
  }
  groundBytes, readErr := _assets.ReadFile(path)

  if readErr != nil {
    log.Fatal("Error while reading image: ", readErr)
  }


  img, _, parseErr := image.Decode(bytes.NewReader(groundBytes))
	if parseErr != nil {
    log.Fatal("Error while parsing image: ",parseErr)
	}

  ebitenImage := ebiten.NewImageFromImage(img)

  ga.images[path] = ebitenImage

  return ebitenImage
}

