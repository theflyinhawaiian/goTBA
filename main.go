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

		levelMap := floorplan.GenerateMap()
		floorplan.Illustrate(levelMap.Start, levelMap.End, levelMap.Grid)

		fmt.Scanln(&input)
	}

	fmt.Println("buh bye!")
}
