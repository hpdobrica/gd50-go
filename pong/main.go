package main

import (
	"fmt"
	"image/color"
	"log"
"math/rand"
"os"
	"strconv"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
  servingPlayer = 1
  winningPlayer int

  gameState = "start"

  justServed bool

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
    fmt.Println("keypress happened", gameState)
    if gameState == "start" {
      gameState = "serve"
    } else if gameState == "serve" {

      if servingPlayer == 1 {
        ball.dx = 100
      } else {
        ball.dx = -100
      }

      ball.dy = float64(rand.Intn(50 - (-50)) + (-50) * 1.5)
      gameState = "play"
      justServed = true
      fmt.Println("setting game state to play")
    } else if gameState == "victory" {

      gameState = "start"
      playerOneScore = 0
      playerTwoScore = 0
      servingPlayer = 1

      ball.Reset()

    }


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

  // enable ai oponnent
  // if playerTwo.y + playerTwo.height/2 < ball.y {
    // playerTwo.dy = paddleSpeed
  // } else if playerTwo.y + playerTwo.height/2 > ball.y  {
    // playerTwo.dy = -paddleSpeed
  // } else {
    // playerTwo.dy = 0
  // }

  // block controls when ball comes to your side of the screen
  if ball.x <= screenWidth/2 && !(servingPlayer == 2 && justServed == true) {
    playerOne.dy = 0
  }
  if ball.x >= screenWidth/2 && !(servingPlayer == 1 && justServed == true) {
    playerTwo.dy = 0
  }



  if gameState == "play" {
    ball.Update()
  }

  playerOne.Update()
  playerTwo.Update()



  if ball.IsColliding(playerOne) {
    ball.x = playerOne.x + playerOne.width
    ball.dx = -ball.dx * 1.03
    justServed = false
    // ball.dy = float64(rand.Intn(50 - (-50)) + (-50) * 1.5)
  }

  if ball.IsColliding(playerTwo) {
    ball.x = playerTwo.x - ball.radius
    ball.dx = -ball.dx * 1.03
    justServed = false
    // ball.dy = float64(rand.Intn(50 - (-50)) + (-50) * 1.5)
  }

  if ball.y < 0 {
    ball.y = 0
    ball.dy = -ball.dy
  }

  if ball.y + ball.radius > screenHeight {
    ball.y = screenHeight - ball.radius
    ball.dy = -ball.dy
  }


  if ball.x < 0 {
    playerOneScore += 1
    servingPlayer = 2
    gameState = "serve"
    ball.Reset()
  }

  if ball.x > screenWidth {
    playerTwoScore += 1
    servingPlayer = 1
    gameState = "serve"
    ball.Reset()
  }

  if playerOneScore >= 5  {
    gameState = "victory"
    winningPlayer = 1
  }

  if playerTwoScore >= 5 {
    gameState = "victory"
    winningPlayer = 2
  }

  return nil

}

func (g *Game) Draw(screen *ebiten.Image) {


  screen.Clear()

  ebitenutil.DebugPrint(screen, strconv.FormatFloat(ebiten.CurrentFPS(), 'f', 1, 32))

  
  // title
	str := gameState

  if gameState == "serve" {
    str = gameState + " " + fmt.Sprint(servingPlayer)
  }

  if gameState == "victory" {
    str = "player " +fmt.Sprint(winningPlayer) + " won"
  }

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
