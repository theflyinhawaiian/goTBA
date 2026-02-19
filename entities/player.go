package entities

type Player struct {
	Stats     Stats
	Inventory Inventory
}

func CreatePlayer() Player {
	return Player{Stats: generatePlayerStats()}
}
