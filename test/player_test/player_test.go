package player_test

import (
	"log"
	"testing"

	"github.com/arne-zillhardt/raycasting/pkg/player"
)

func TestPlayerCreation(t *testing.T) {
	log.Println("TEST: player creation test")
	player := player.GetInstance()

	if player.GetPositionX() != 1 {
		log.Println("Start-PositionX not 1")
		t.Fail()
	}
	if player.GetPositionY() != 1 {
		log.Println("Start-PositionY not 1")
		t.Fail()
	}
}

func TestPlayerMoveForward(t *testing.T) {
	log.Println("TEST: player creation test")
	player := player.GetInstance()

	player.Move(player.MovementDirection.FORWARD)
}
