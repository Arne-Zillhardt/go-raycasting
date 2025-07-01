package game_test

import (
	"log"
	"testing"

	"github.com/arne-zillhardt/raycasting/pkg/game"
)

func TestInitialization(t *testing.T) {
	log.Println("TEST: Initialization of game with gameMap")

	gameWithGameMap := game.GetInstance()

	for val1 := range(10) {
		if !gameWithGameMap.IsWall(uint32(0), uint32(val1)) {
			log.Println("No first horizontal wall")
			log.Println("X: 0 Y: ", val1)
			t.FailNow()
		}
		if !gameWithGameMap.IsWall(uint32(9), uint32(val1)) {
			log.Println("No last horizontal wall")
			log.Println("X: 9 Y: ", val1)
			t.FailNow()
		}
		if !gameWithGameMap.IsWall(uint32(val1), uint32(0)) {
			log.Println("No last vertical wall")
			log.Println("X: ", val1, " Y: 0")
			t.FailNow()
		}
		if !gameWithGameMap.IsWall(uint32(val1), uint32(9)) {
			log.Println("No first vertical wall")
			log.Println("X: ", val1, " Y: 9")
			t.FailNow()
		}
	//if !foundWall {
	//	log.Println("Should be atleas one wall per row")
	//	t.FailNow()
	//}
	}
}
