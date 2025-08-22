package pkg

type MovementDirections struct {
	
}

type player struct {
	positionX uint
	positionY uint
}

func createPlayer(gameMap [][]uint) (player player) {
	for y := range(len(gameMap)) {
		for x := range(len(gameMap[y])) {
			if x == 0 || y == 0 {
				continue
			}

			if gameMap[y][x] == 0 {
				player.positionX = uint(x)
				player.positionY = uint(y)
				return player
			}
		}
	}

	return player
}
