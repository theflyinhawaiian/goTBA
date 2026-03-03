package entities

type Enemy struct {
	Name      string
	maxHealth int
	attack    int
	defense   int
	hitChance float32
}

func CreateHobGoblin() Enemy {
	return Enemy{
		Name:      "Hobgoblin",
		maxHealth: 50,
		attack:    10,
		defense:   3,
		hitChance: 0.8,
	}
}

func CreateGoblin() Enemy {
	return Enemy{
		Name:      "Goblin",
		maxHealth: 20,
		attack:    5,
		defense:   1,
		hitChance: 0.95,
	}
}

func CreateRat() Enemy {
	return Enemy{
		Name:      "Rat",
		maxHealth: 10,
		attack:    2,
		defense:   0,
		hitChance: 0.66,
	}
}
