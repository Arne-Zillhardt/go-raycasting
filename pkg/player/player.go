package player

import (
	"log"
)

type Player interface {
}
type player struct {
}

func GetInstance() Player {
	log.Println("created player")
	return player{}
}
