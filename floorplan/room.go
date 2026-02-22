package floorplan

type Room struct {
	exists      bool
	connections int
	nums        int
}

func createEmptyRoom() Room {
	return Room{exists: true}
}
