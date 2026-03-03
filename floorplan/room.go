package floorplan

import "tba/entities"

type Room struct {
	exists      bool
	connections int
	Enemies     []entities.Enemy
	Items       []entities.Item
}

func createEmptyRoom() Room {
	return Room{exists: true}
}
