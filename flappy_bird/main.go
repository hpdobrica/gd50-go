package main

import (
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var assets GameAssets

const windowWidth = 1280
const windowHeight = 720

const screenWidth = 512
const screenHeight = 288

const backgroundSpeed = 30
const groundSpeed = 60
const backgroundLoopingPoint = 413

var groundOffset float64
var backgroundOffset float64

const dt float64 = 1 / 60.0

var scrolling = true

var bird Bird

var pipeManager PipeManager

type Game struct{}

func (g Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if scrolling {
		groundOffset = math.Mod(groundOffset+(groundSpeed*dt), screenWidth)
		backgroundOffset = math.Mod(backgroundOffset+(backgroundSpeed*dt), backgroundLoopingPoint)

		bird.Update()

		for _, pipePair := range pipeManager.pipePairs {
			pipePair.Update()
			if bird.Collides(*pipePair.top) || bird.Collides(*pipePair.bottom) {
				scrolling = false
			}
		}

		ground := assets.getImage("assets/ground.png")
		_, groundHeight := ground.Size()
		if bird.y+float64(bird.height) > float64(screenHeight-groundHeight) {
			scrolling = false
		}

		pipeManager.ManageLifecycle()

	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {

	drawBackground(screen)
	for _, pipePair := range pipeManager.pipePairs {
		pipePair.Draw(screen)
	}
	drawGround(screen)

	bird.Draw(screen)

}

func main() {
	rand.Seed(time.Now().UnixNano())

	assets = NewGameAssets()

	bird = NewBird()
	pipeManager = NewPipeManager()

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Flappy Bird!")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}

func drawGround(screen *ebiten.Image) {
	ground := assets.getImage("assets/ground.png")
	_, groundHeight := ground.Size()

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(-groundOffset, float64(screenHeight-groundHeight))
	screen.DrawImage(ground, opts)

}

func drawBackground(screen *ebiten.Image) {
	background := assets.getImage("assets/background.png")
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(-backgroundOffset, 0)
	screen.DrawImage(background, opts)

}
