package main

import (
	"log"

	"github.com/sflewis2970/ninja-game/game"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	game.StartGame()
}
