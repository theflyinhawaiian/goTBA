package main

import (
	"fmt"
	"tba/entities"
)

func main() {
	input := ""

	for input != "q" {
		stats := entities.GeneratePlayerStats()
		fmt.Printf("Your Stats:\nHealth: %d\nAttack: %d\nDefense: %d\nSpeed: %d\nMagic: %d\n",
			stats.Health,
			stats.Attack,
			stats.Defense,
			stats.Speed,
			stats.Magic)

		fmt.Println("Press enter or q to quit")

		fmt.Scanln(&input)
	}

	fmt.Println("buh bye!")
}
