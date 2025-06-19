package player_test

import (
	"log"
	"testing"

	"github.com/arne-zillhardt/raycasting/pkg/player"
)

func TestPlayerCreation(t *testing.T) {
	log.Println("TEST: player creation test")
	player := player.GetInstance()

	t.Fail()
}
