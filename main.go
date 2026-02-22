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
		floorplan.Illustrate(floorplan.Point{X: 100, Y: 100}, floorplan.Point{X: 100, Y: 100}, levelMap)

		fmt.Scanln(&input)
	}

	fmt.Println("buh bye!")
}
