package pkg

import (
	"log"
	"math/rand"
)

var chanceOfWall float64 = 0.25

func getGameMap(width uint, height uint) (gameMap [][]uint) {
	gameMap = make([][]uint, width)

	for y := range(height) {
		var line []uint = make([]uint, width)

		for x := range(width) {
			inFirstLine := y == 0
			inLastLine := y == height - 1
			inFirstField := x == 0
			inLastField := x == width - 1
			wall := rand.Float64() >= (1.0 - chanceOfWall)

			if inFirstLine || inLastLine || inFirstField || inLastField || wall {
				line[x] = uint(1)
			}

			log.Println(line[x])
		}

		gameMap[y] = line
		log.Println()
	}

	return gameMap
}
