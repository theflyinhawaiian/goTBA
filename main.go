package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	fp "tba/floorplan"
)

func main() {
	input := ""
	levelMap := fp.GenerateMap()
	playerPosition := fp.Point{X: levelMap.Start.X, Y: levelMap.Start.Y}

	for input != "q" {
		// player := entities.CreatePlayer()
		// stats := player.Stats
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fp.Illustrate(playerPosition, levelMap.End, levelMap.Grid)

		exits := fp.GetExitDirections(playerPosition.X, playerPosition.Y, levelMap.Grid)
		if len(exits) == 1 {
			fmt.Printf("There's an exit to the %s\n", GetDirectionDescriptions(exits))
		} else {
			fmt.Printf("There are exits to the %s\n", GetDirectionDescriptions(exits))
		}

		fmt.Println("What do?? ")
		fmt.Scanln(&input)

		dir := choiceStrToDirection(input)

		if !slices.Contains(exits, dir) {
			continue
		}

		switch dir {
		case fp.North:
			playerPosition = fp.Point{X: playerPosition.X, Y: playerPosition.Y + 1}
		case fp.South:
			playerPosition = fp.Point{X: playerPosition.X, Y: playerPosition.Y - 1}
		case fp.East:
			playerPosition = fp.Point{X: playerPosition.X + 1, Y: playerPosition.Y}
		case fp.West:
			playerPosition = fp.Point{X: playerPosition.X - 1, Y: playerPosition.Y}
		}
	}

	fmt.Println("buh bye!")
}

func choiceStrToDirection(rawChoice string) fp.Direction {
	choice := strings.ToLower(rawChoice)
	switch choice {
	case "n":
		fallthrough
	case "north":
		return fp.North
	case "s":
		fallthrough
	case "south":
		return fp.South
	case "e":
		fallthrough
	case "east":
		return fp.East
	case "w":
		fallthrough
	case "west":
		return fp.West
	}

	return -1
}

func GetDirectionDescriptions(directions []fp.Direction) string {
	directionText := make([]string, 0)
	for _, direction := range directions {
		switch direction {
		case fp.North:
			directionText = append(directionText, "North")
		case fp.South:
			directionText = append(directionText, "South")
		case fp.East:
			directionText = append(directionText, "East")
		case fp.West:
			directionText = append(directionText, "West")
		}
	}

	switch len(directionText) {
	case 1:
		return directionText[0]
	case 2:
		return fmt.Sprintf("%s and %s", directionText[0], directionText[1])
	case 3:
		return fmt.Sprintf("%s, %s, and %s", directionText[0], directionText[1], directionText[2])
	case 4:
		return fmt.Sprintf("%s, %s, %s, and %s", directionText[0], directionText[1], directionText[2], directionText[3])
	default:
		panic("Ahhhhhh there are either zero or more than four exits, what is happening")
	}

}
