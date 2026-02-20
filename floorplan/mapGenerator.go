package floorplan

import (
	"fmt"
	"math/rand"
)

type point struct {
	x int
	y int
}

var offsets = []int{-1, 0, 1}

func GenerateMap() Room {
	grid := make([][]Room, 31)
	for i := range grid {
		grid[i] = make([]Room, 31)
	}

	roomCount := 0
	numRooms := rand.Intn(4) + 12
	deadEndRooms := make([]point, 1)
	currX := 16
	currY := 16
	firstRoom := point{x: currX, y: currY}
	grid[currX][currY] = createEmptyRoom()
	deadEndRooms[0] = firstRoom

	for roomCount < numRooms {
		targetRoom := deadEndRooms[roomCount]
		possibleExits := getLegalExits(targetRoom.x, targetRoom.y, grid)
		generatedRoom := false
		for _, exit := range possibleExits {
			if rand.Intn(10)+1 > 5 {
				newRoom := createEmptyRoom()
				roomCount++
				grid[exit.x][exit.y] = newRoom

				idx := getExitIndex(targetRoom, exit)
				inverseIdx := getExitIndex(exit, targetRoom)
				grid[targetRoom.x][targetRoom.y].exits[idx] = newRoom
				grid[exit.x][exit.y].exits[inverseIdx] = grid[targetRoom.x][targetRoom.y]

				deadEndRooms = append(deadEndRooms, exit)
				generatedRoom = true
			}
		}

		if !generatedRoom {
			deadEndRooms = append(deadEndRooms, targetRoom)
		}
	}

	fmt.Printf("Successfully generated %d rooms\n", roomCount)

	output := ""
	deadEndCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch {
			case len(grid[i][j].exits) == 1:
				deadEndCount++
				output += "X"
			case grid[i][j].exits != nil:
				output += "O"
			default:
				output += "-"
			}
		}
		output += "\n"
	}

	fmt.Printf("There are %d dead end rooms on the map\n", deadEndCount)
	fmt.Println(output)

	return grid[currX][currY]
}

func getExitIndex(p1 point, p2 point) int {
	switch {
	case p1.x-p2.x == -1 && p1.y == p2.y:
		return 0
	case p1.x == p2.x && p1.y-p2.y == 1:
		return 1
	case p1.x-p2.x == 1 && p1.y == p2.y:
		return 2
	case p1.x == p2.x && p1.y-p2.y == -1:
		return 3
	default:
		panic("points are not neighbors")
	}

}

func getLegalExits(x int, y int, grid [][]Room) []point {
	exits := make([]point, 0)
	for _, i := range offsets {
		for _, j := range offsets {
			// check cardinal directions only!
			if (i == 0 || j == 0) && i != j {
				nonNilNeighbors := getNonNilNeighbors(x+i, y+j, grid)
				if nonNilNeighbors == 1 && grid[x+i][y+j].exits == nil {
					exits = append(exits, point{x: x + i, y: y + j})
				}
			}
		}
	}

	return exits
}

func getNonNilNeighbors(x int, y int, grid [][]Room) int {
	numNeighbors := 0
	for _, i := range offsets {
		for _, j := range offsets {
			// check cardinal directions only!
			if (i == 0 || j == 0) && i != j {
				if grid[x+i][y+j].exits != nil {
					numNeighbors++
				}
			}
		}
	}

	return numNeighbors
}
