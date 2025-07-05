package game

import (
	"log"
	"math"
)

type player struct {
	positonX  *int32
	positonY  *int32
	viewAngle *float64
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
		positonX:  new(int32),
		positonY:  new(int32),
		viewAngle: new(float64),
	}

	*player.positonX = int32(1)
	*player.positonY = int32(1)
	*player.viewAngle = float64(90)
	return player
}

func (p player) move(direction movementDirection) {
	log.Println("Moving the player ", direction)
	log.Println("Before moving ", *p.positonX, " ", *p.positonY)

	*p.positonX += int32(direction.additionToPositionX)
	*p.positonY += int32(direction.additionToPositionY)

	log.Println("After moving ", *p.positonX, " ", *p.positonY)
}

func (p player) changeView(change viewChange) {
	*p.viewAngle += float64(change)
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

func (p player) calculateLinePosition() (float64, float64) {
	length := 5
	angleInRadiens := *p.viewAngle * (math.Pi/180)

	lineX := float64(*p.positonX) + float64(length) * math.Sin(angleInRadiens)
	lineY := float64(*p.positonY) + float64(length) * math.Cos(angleInRadiens)

	log.Println("Calculated x", lineX)
	log.Println("Calculated y", lineY)

	return lineX, lineY
}
