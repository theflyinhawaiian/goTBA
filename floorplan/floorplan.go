package floorplan

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type Map struct {
	Start Point
	End   Point
	Grid  [][]Room
}
