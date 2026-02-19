package main

import "fmt"

func main() {
	input := ""

	for input != "q" {
		stats := generatePlayerStats()
		fmt.Printf("Your Stats:\nHealth: %d\nAttack: %d\nDefense: %d\nSpeed: %d\nMagic: %d\n",
			stats.health,
			stats.attack,
			stats.defense,
			stats.speed,
			stats.magic)

		fmt.Println("Press enter or q to quit")

		fmt.Scanln(&input)
	}

	fmt.Println("buh bye!")
}
