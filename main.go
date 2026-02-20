package main

import (
	"fmt"
	"tba/floorplan"
)

func main() {
	input := ""

	for input != "q" {
		// player := entities.CreatePlayer()
		// stats := player.Stats

		_ = floorplan.GenerateMap()

		fmt.Scanln(&input)
	}

	fmt.Println("buh bye!")
}
