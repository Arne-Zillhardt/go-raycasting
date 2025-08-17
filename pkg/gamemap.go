package pkg

import (
	"log"
)

func getGameMap(width uint, height uint) (gameMap [][]uint) {
	gameMap = make([][]uint, width)

	for y := range(height) {
		var line []uint = make([]uint, width)

		for x := range(width) {
			log.Println(line)
			log.Println(x)
		}

		gameMap[y] = line
	}

	return gameMap
}
