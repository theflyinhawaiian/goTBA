package floorplan

type Room struct {
	exits map[int]Room
	nums  int
}

func createEmptyRoom() Room {
	return Room{exits: make(map[int]Room)}
}
