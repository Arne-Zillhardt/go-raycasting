package pkg

const planeX float64 = 0
const planeY float64 = 0.66

type player struct {
	positionX float64
	positionY float64
	dirX float64
	dirY float64
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

func (p *player) getRayDir(horizontalPosition int, horizontalLength int) (float64, float64) {
	var positionOnCamera float64 = 2 * float64(horizontalPosition) / float64(horizontalLength) - 1

	rayDirX := p.dirX + planeX * positionOnCamera
	rayDirY := p.dirY + planeY * positionOnCamera

	return rayDirX, rayDirY
}
