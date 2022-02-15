package main

import (
	"math/rand"
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const windowWidth = 1024
const windowHeight = 768
const scale = 2

const screenWidth = windowWidth/scale
const screenHeight = windowHeight/scale
const fontSizeSmall = 8
const fontSizeLarge = 30
const dpi = 72

const paddleSpeed = 200
const dt float64 = 1/60.0




var (
	arcadeFontSmall font.Face
	arcadeFontLarge font.Face

  playerOneScore = 0
  playerTwoScore = 0

  gameState = "start"

  playerOne Paddle
  playerTwo Paddle
  ball Ball
)

type Game struct{}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
    os.Exit(0)
	}
  
  if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
    if gameState == "start" {
      gameState = "play"
    } else {
      gameState = "start"

      ball.Reset()

    }

    if rand.Intn(10 - 1) + 1 > 5 {
      ball.dx = 100 
    } else {
      ball.dx = -100
    }
    ball.dy = float64(rand.Intn(50 - (-50)) + (-50) * 1.5)

  }


	if ebiten.IsKeyPressed(ebiten.KeyW) {
    playerOne.dy = -paddleSpeed
  } else if ebiten.IsKeyPressed(ebiten.KeyS) {
    playerOne.dy = paddleSpeed
  } else {
    playerOne.dy = 0
  }

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
    playerTwo.dy = -paddleSpeed
  } else if ebiten.IsKeyPressed(ebiten.KeyDown) {
    playerTwo.dy = paddleSpeed
  } else {
    playerTwo.dy = 0
  }


  if gameState == "play" {
    ball.Update()
  }

  playerOne.Update()
  playerTwo.Update()



  if ball.IsColliding(playerOne) {
    ball.x = playerOne.x + playerOne.width
    ball.dx = -ball.dx * 1.2
    // ball.dy = float64(rand.Intn(50 - (-50)) + (-50) * 1.5)
  }

  if ball.IsColliding(playerTwo) {
    ball.x = playerTwo.x - ball.radius
    ball.dx = -ball.dx * 1.2
    // ball.dy = float64(rand.Intn(50 - (-50)) + (-50) * 1.5)
  }

  if ball.y < 0 {
    ball.y = 0
    ball.dy = -ball.dy
    fmt.Println("switch1")
  }

  if ball.y + ball.radius > screenHeight {
    ball.y = screenHeight - ball.radius
    ball.dy = -ball.dy
    fmt.Println("switch2")
  }


  return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
  /* fmt.Println("fps", ebiten.CurrentFPS())
  fmt.Println("tps", ebiten.CurrentTPS()) */


  screen.Clear()

  // title
	str := "Hello " + gameState
	x := (screenWidth - len(str)*fontSizeSmall) / 2
	text.Draw(screen, str, arcadeFontSmall, x, 20, color.White)

  // player one
  playerOne.Draw(screen)

  text.Draw(screen, fmt.Sprint(playerOneScore), arcadeFontLarge, 60+fontSizeLarge, 60, color.White)

  // player two
  playerTwo.Draw(screen)

  text.Draw(screen, fmt.Sprint(playerTwoScore), arcadeFontLarge, screenHeight-60+fontSizeLarge, 60, color.White)

  // ball
  ball.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
  rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Hello, Pong!")

	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	arcadeFontSmall, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSizeSmall,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	arcadeFontLarge, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSizeLarge,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})


  playerOne = NewPaddle(0, 10, 5, 50)
  playerTwo = NewPaddle(screenWidth-5, screenHeight-60, 5, 50)
  ball = NewBall(screenWidth/2-2, screenHeight/2-2, 4)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
