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
		playerPosition := floorplan.Point{X: levelMap.Start.X, Y: levelMap.Start.Y}
		floorplan.Illustrate(levelMap.Start, levelMap.End, levelMap.Grid)

		exits := floorplan.GetNeighborOffsets(playerPosition.X, playerPosition.Y, levelMap.Grid)
		fmt.Printf("Exits are %s", GetExitDescription(exits))
		var choice string

		fmt.Scanln(&choice)

		fmt.Scanln(&input)
	}

	fmt.Println("buh bye!")
}

func GetExitDescription(exits []floorplan.Point) string {
	exitText := make([]string, 0)
	for _, exit := range exits {
		switch {
		case exit.Y == 1:
			exitText = append(exitText, "North")
		case exit.Y == -1:
			exitText = append(exitText, "South")
		case exit.X == 1:
			exitText = append(exitText, "East")
		case exit.X == -1:
			exitText = append(exitText, "West")
		}
	}

	switch len(exitText) {
	case 1:
		return exitText[0]
	case 2:
		return fmt.Sprintf("%s and %s", exitText[0], exitText[1])
	case 3:
		return fmt.Sprintf("%s, %s, and %s", exitText[0], exitText[1], exitText[2])
	case 4:
		return fmt.Sprintf("%s, %s, %s, and %s", exitText[0], exitText[1], exitText[2], exitText[3])
	default:
		panic("Ahhhhhh there are either zero or more than four exits, what is happening")
	}

}
