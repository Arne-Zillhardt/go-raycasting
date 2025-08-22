package main

import (
	"log"

	"github.com/arne-zillhardt/raycasting/pkg"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	application := &pkg.Application{}
	application.Init()

	if err := ebiten.RunGame(application); err != nil {
		log.Fatal(err)
	}
}
