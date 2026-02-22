package floorplan

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
)

type Point struct {
	X int
	Y int
}

var offsets = []int{-1, 0, 1}

func GenerateMap() [][]Room {
	// generate a 31x31 grid so we can guarantee we won't go out of grid bounds when generating
	grid := make([][]Room, 31)
	for i := range grid {
		grid[i] = make([]Room, 31)
	}

	roomCount := 1
	// target is 12-15 rooms
	roomNumTarget := rand.Intn(4) + 12
	deadEndRooms := make([]Point, 1)

	// Generate our first room in the middle and set it as the first dead end room
	currX := 16
	currY := 16
	firstRoom := Point{X: currX, Y: currY}
	grid[currX][currY] = createEmptyRoom()
	deadEndRooms[0] = firstRoom
	deadEndIdx := 0

	// While we don't have enough rooms, grab the next dead end room in the slice and iterate
	// through all the neighbors
StartGeneration:
	for roomCount < roomNumTarget {
		targetRoom := deadEndRooms[deadEndIdx]
		possibleExits := getLegalExits(targetRoom.X, targetRoom.Y, grid)
		generatedRoom := false
		for _, exit := range possibleExits {
			if roomCount >= roomNumTarget {
				break StartGeneration
			}
			if rand.Intn(10)+1 > 5 {
				grid[exit.X][exit.Y] = createEmptyRoom()
				roomCount++

				// associate the two rooms with one another
				grid[targetRoom.X][targetRoom.Y].connections++
				grid[exit.X][exit.Y].connections++

				deadEndRooms = append(deadEndRooms, exit)
				generatedRoom = true
			}
		}

		if !generatedRoom && len(possibleExits) > 0 {
			deadEndRooms = append(deadEndRooms, targetRoom)
		}

		deadEndIdx++
	}

	min := Point{math.MaxInt, math.MaxInt}
	max := Point{}
	deadEndRooms = deadEndRooms[:0]
	for i := range grid {
		for j := range grid[i] {
			if !grid[i][j].exists {
				continue
			}
			if i > max.X {
				max.X = i
			}
			if j > max.Y {
				max.Y = j
			}
			if i < min.X {
				min.X = i
			}
			if j < min.Y {
				min.Y = j
			}
			if grid[i][j].connections == 1 {
				deadEndRooms = append(deadEndRooms, Point{X: i, Y: j})
			}
		}
	}

	floorGrid := make([][]Room, max.X-min.X+1)
	for i := range floorGrid {
		floorGrid[i] = make([]Room, max.Y-min.Y+1)
	}

	for i := range floorGrid {
		for j := range floorGrid[i] {
			floorGrid[i][j] = grid[min.X+i][min.Y+j]
		}
	}

	return floorGrid
}

func cls() {
	cmd := exec.Command("cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Illustrate(curr Point, new Point, grid [][]Room) {
	cls()
	output := ""
	deadEndCount := 0
	roomCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].exists {
				roomCount++
			}
			switch {
			case i == new.X && j == new.Y:
				output += "*"
			case i == curr.X && j == curr.Y:
				output += "^"
			case grid[i][j].connections == 1:
				deadEndCount++
				output += "X"
			case grid[i][j].exists:
				output += "O"
			default:
				output += "-"
			}
		}
		output += "\n"
	}

	fmt.Println(output)
}

func getLegalExits(x int, y int, grid [][]Room) []Point {
	exits := make([]Point, 0)
	for _, i := range offsets {
		for _, j := range offsets {
			// check cardinal directions only!
			if (i == 0 || j == 0) && i != j {
				neighborCount := getExistingNeighborCount(x+i, y+j, grid)
				if neighborCount == 1 && !grid[x+i][y+j].exists {
					exits = append(exits, Point{X: x + i, Y: y + j})
				}
			}
		}
	}

	return exits
}

func getExistingNeighborCount(x int, y int, grid [][]Room) int {
	numNeighbors := 0
	for _, i := range offsets {
		for _, j := range offsets {
			// check cardinal directions only!
			if (i == 0 || j == 0) && i != j {
				if grid[x+i][y+j].exists {
					numNeighbors++
				}
			}
		}
	}

	return numNeighbors
}
