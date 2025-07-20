package game

import (
	"image/color"
	"log"
	"math"
	"time"

	//"github.com/arne-zillhardt/raycasting/pkg/raycasting"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game interface {
	Update() error
	Draw(*ebiten.Image)
	Layout(int, int) (int, int)
	PlayerPosition() (int, int)
	IsWall(uint32, uint32) bool
	Size() int
}
type game struct {
	playingField gameMap
	player       player
}

var w int = 640
var h int = 480
var oldTime time.Time

func (g game) Update() error {
	currentTime := time.Now()
	var deltaTime float64 = float64(currentTime.UnixNano() - oldTime.UnixNano())/1000000000

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.move(FORWARD, deltaTime)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.move(BACKWARD, deltaTime)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.move(LEFT, deltaTime)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.move(RIGHT, deltaTime)
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.player.changeView(INCREASE, deltaTime)
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.player.changeView(DECREASE, deltaTime)
	}

	oldTime = time.Now()
	return nil
}

func (g game) Draw(screen *ebiten.Image) {
	imageToDraw := ebiten.NewImage(w, h)

	drawFloorAndCeiling(imageToDraw)

	var dirX float64 = 1
	var dirY float64 = 0
  var planeX float64 = 0
	var planeY float64 = 0.66

	for x := range(w) {
		cameraX := 2 * float64(x) / float64(w) - 1
		vectorX := dirX + planeX * cameraX
		vectorY := dirY + planeY * cameraX;

		deltaDistX := math.Abs(1 / vectorX)
		if vectorX == 0 {
			deltaDistX = 1e30 
		}
		deltaDistY := math.Abs(1 / vectorY)
		if vectorY == 0 {
			deltaDistY = 1e30 
		}

		var stepX int
		var stepY int

		var sideDistX float64
		var sideDistY float64

		mapX, mapY := g.player.position()
		posX := *g.player.positonX
		posY := *g.player.positonY

		if vectorX < 0 {
			stepX = -1
			sideDistX = (posX - float64(mapX)) * deltaDistX
		} else {
			stepX = 1
			sideDistX = (float64(mapX )+ 1.0 - posX) * deltaDistX
		}
		if vectorY < 0 {
			stepY = -1
			sideDistY = (posY - float64(mapY)) * deltaDistY
		} else {
			stepY = 1;
			sideDistY = (float64(mapY )+ 1.0 - posY) * deltaDistY
		}

		side := 0
		hit := 0
		for hit == 0 {
			if sideDistX < sideDistY {
				sideDistX += deltaDistX
				mapX += stepX
				side = 0
			} else {
				sideDistY += deltaDistY
				mapY += stepY
				side = 1
			}

			if g.playingField.playMap[mapY][mapX] > 0 {
				hit = 1
			}
		}

		perpWallDist := (sideDistY - deltaDistY);
		if side == 0 {
			perpWallDist = (sideDistX - deltaDistX);
		}

		lineHeight := int(float64(h) / perpWallDist)

		drawStart := -lineHeight / 2 + h / 2
		if drawStart < 0 {
			drawStart = 0
		}
		
		drawEnd := lineHeight / 2 + h / 2
		if drawEnd >= h {
			drawEnd = h - 1
		}

		colorToDraw := color.RGBA{0xff, 0xff, 0xff, 0xff}
		if side == 1 {
			colorToDraw = color.RGBA{0x80, 0x80, 0x80, 0xff}
		}
		vector.StrokeLine(imageToDraw, float32(x), float32(drawStart), float32(x), float32(drawEnd), 1, colorToDraw, false)
	}

	drawMap(imageToDraw, g)
	screen.DrawImage(imageToDraw, nil)
}

func drawFloorAndCeiling(imageToDraw *ebiten.Image) {

	colorToDraw := color.RGBA{0x3C, 0x3C, 0xE0, 0xff}
	for y := range h {
		for x := range w {
			colorToDraw = color.RGBA{0x70, 0x70, 0x70, 0xff}

			if y <= h/2 {
				colorToDraw = color.RGBA{0x38, 0x38, 0x38, 0xff}
			}

			imageToDraw.Set(x, y, colorToDraw)
		}
	}
}

func drawMap(imageToDraw *ebiten.Image, g game) {
	playerPositionX, playerPositionY := g.player.position()
	fieldSize := g.playingField.size()
	for y := range fieldSize {
		for x := range fieldSize {
			colorToDraw := color.RGBA{0x3C, 0x3C, 0xE0, 0xff}
			if !g.playingField.isWall(uint32(x), uint32(y)) {
				colorToDraw = color.RGBA{0xB5, 0x9B, 0x41, 0xff}
			}
			
			if x == playerPositionX && y == playerPositionY {
				colorToDraw = color.RGBA{0xff, 0x00, 0x00, 0xff}
			}

			imageToDraw.Set(x, y, colorToDraw)
		}
	}
}

func (g game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return w, h
}

func GetInstance() Game {
	log.Println("Created game")
	return game{
		player:       initPlayer(),
		playingField: initGameMap(),
	}
}

func (g game) PlayerPosition() (int, int) {
	return g.player.position()
}

func (g game) IsWall(positionX uint32, positionY uint32) bool {
	return g.playingField.isWall(positionX, positionY)
}

func (g game) Size() int {
	return g.playingField.size()
}
