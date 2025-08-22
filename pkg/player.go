package pkg

type MovementDirections struct {
}

type player struct {
	positionX float64
	positionY float64
}

func createPlayer(gameMap [][]uint) (player player) {
	for y := range len(gameMap) {
		for x := range len(gameMap[y]) {
			if x == 0 || y == 0 {
				continue
			}

			if gameMap[y][x] == 0 {
				player.positionX = float64(x)
				player.positionY = float64(y)
				return player
			}
		}
	}

	return player
}
