package player

import (
	"log"
)

type Player interface {
	Move(MovementDirection)
	GetPositionX() int32
	GetPositionY() int32
}
type player struct {
	positonX int32
	positonY int32
}

type MovementDirection int16

const (
	FORWARD MovementDirection = 1
)

func GetInstance() Player {
	log.Println("created player")
	return player{
		positonX: int32(1),
		positonY: int32(1),
	}
}

func (p player) Move(direction MovementDirection) {

}

func (p player) GetPositionX() int32 {
	return p.positonX

}

func (p player) GetPositionY() int32 {
	return p.positonY
}
