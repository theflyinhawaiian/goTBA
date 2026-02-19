package main

import (
	"math/rand"
)

type Stats struct {
	health  int
	attack  int
	defense int
	speed   int
	magic   int
}

func generatePlayerStats() Stats {
	// generate 5 random numbers, capping the sum at 41
	statsRaw := make([]int, 5)
	availablePoints := 25
	for i := range len(statsRaw) - 1 {
		// max stat is either 10 or the largest number that would leave enough
		// for the rest of the stats to be 1, whichever is smaller
		remainingStats := len(statsRaw) - i - 1
		maxStat := min(10, availablePoints-(remainingStats))
		minStat := max(1, availablePoints-(remainingStats)*10)
		stat := rand.Intn(maxStat-minStat+1) + minStat
		statsRaw[i] = stat
		availablePoints -= stat
	}

	// assign the last stat to be whatever number is leftover
	statsRaw[len(statsRaw)-1] = availablePoints

	// permute these because numbers generated earlier will tend to be larger
	// and we want these stats to be randomized
	playerStats := make([]int, 5)
	for i, v := range rand.Perm(len(statsRaw)) {
		playerStats[v] = statsRaw[i]
	}

	return Stats{
		health:  playerStats[0],
		attack:  playerStats[1],
		defense: playerStats[2],
		speed:   playerStats[3],
		magic:   playerStats[4],
	}
}
