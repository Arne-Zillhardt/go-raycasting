package gui

import (
	"log"

	"github.com/arne-zillhardt/raycasting/pkg/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func RunGame() {
	ebiten.SetWindowTitle("Raycasting")
	if err := ebiten.RunGame(game.GetInstance()); err != nil {
		log.Fatal("err")
	}
}
