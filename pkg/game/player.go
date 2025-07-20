package game

import (
	"log"
	"math"
)

type player struct {
	positonX  *float64
	positonY  *float64
	viewAngle *float64
	dirX *float64
	dirY * float64
}

type movementDirection struct {
	additionToPositionX int16
	additionToPositionY int16
}
type viewChange int16

var (
	FORWARD  movementDirection = movementDirection{0, -1}
	BACKWARD movementDirection = movementDirection{0, 1}
	LEFT     movementDirection = movementDirection{-1, 0}
	RIGHT    movementDirection = movementDirection{1, 0}

	INCREASE viewChange = 5
	DECREASE viewChange = -5
)

func initPlayer() player {
	log.Println("Created player")
	player := player{
		positonX:  new(float64),
		positonY:  new(float64),
		viewAngle: new(float64),
		dirX: new(float64),
		dirY: new(float64),
	}

	*player.positonX = float64(1)
	*player.positonY = float64(1)
	*player.viewAngle = float64(90)
	*player.dirX = float64(-1)
	*player.dirY = float64(0)
	return player
}

func (p player) move(direction movementDirection, deltaTime float64) {
	log.Println("Moving the player ", direction)
	log.Println("Before moving ", *p.positonX, " ", *p.positonY)

	*p.positonX += float64(direction.additionToPositionX) * deltaTime
	*p.positonY += float64(direction.additionToPositionY) * deltaTime

	log.Println("After moving ", *p.positonX, " ", *p.positonY)
}

func (p player) changeView(change viewChange, deltaTime float64) {
	*p.viewAngle += float64(change) * deltaTime
	if *p.viewAngle < 0 {
		*p.viewAngle = 360 + *p.viewAngle
	}

	if *p.viewAngle > 360 {
		*p.viewAngle = *p.viewAngle - 360
	}

	log.Println("ViewAngle: ", *p.viewAngle)
}

func (p player) position() (int, int) {
	return int(*p.positonX), int(*p.positonY)
}

func (p player) calculateLinePosition(changeOfAngle float64) (float64, float64) {
	length := 1
	playerAngle := (*p.viewAngle + changeOfAngle)
	log.Println("Angle to calculate with: ", playerAngle)
	angleInRadiens :=  playerAngle * (math.Pi/180)

	lineX := float64(*p.positonX) + float64(length) * math.Sin(angleInRadiens)
	lineY := float64(*p.positonY) + float64(length) * math.Cos(angleInRadiens)

	log.Println("Calculated x", lineX)
	log.Println("Calculated y", lineY)

	return lineX, lineY
}
