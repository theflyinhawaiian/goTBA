package main

import (
	"fmt"
	"tba/entities"
	fp "tba/floorplan"
)

type GameState int

const (
	Exploring GameState = iota
	Combat
	ManagingInventory
	Dead
	Won
)

func (g GameState) String() string {
	switch g {
	case Exploring:
		return "exploring"
	case Combat:
		return "combat"
	case ManagingInventory:
		return "inventory management"
	case Dead:
		return "dead"
	case Won:
		return "won"
	default:
		panic("improper game state toString")
	}
}

type GameEvent struct {
	Type        GameState
	Description string
	Choices     []PlayerChoice
}

type PlayerChoice struct {
	Text                 string
	ValidRepresentations []string
}

var level fp.Map
var state GameState
var prevState GameState
var player entities.Player

func Start(input chan PlayerChoice) <-chan GameEvent {
	level = fp.GenerateLevel()
	player = entities.CreatePlayer()
	updateState(Exploring)

	events := make(chan GameEvent)
	go func() {
		defer close(events)
		var newEvent GameEvent

		events <- GameEvent{Type: Exploring, Choices: []PlayerChoice{{Text: "inventory"}, {Text: "explore"}, {Text: "combat"}}}

		for choice := range input {
			fmt.Printf("player choice: %s\n", choice.Text)
			switch state {
			case Exploring:
				newEvent = processExplorationChoice(choice.Text)
			case Combat:
				newEvent = processCombatChoice(choice.Text)
			case ManagingInventory:
				newEvent = processInventoryChoice(choice.Text)
			}

			var hold string
			fmt.Scanln(&hold)

			updateState(newEvent.Type)

			events <- newEvent
		}

	}()

	return events
}

func processExplorationChoice(choice string) GameEvent {
	switch choice {
	case "explore":
		return GameEvent{Type: Exploring, Choices: []PlayerChoice{{Text: "inventory"}, {Text: "explore"}, {Text: "combat"}}}
	case "combat":
		return GameEvent{Type: Combat, Choices: []PlayerChoice{{Text: "inventory"}, {Text: "explore"}, {Text: "fight"}}}
	case "inventory":
		return GameEvent{Type: ManagingInventory, Choices: []PlayerChoice{{Text: "back"}, {Text: "equip"}}}
	default:
		panic("Bad exploration choice")
	}
}

func processCombatChoice(choice string) GameEvent {
	switch choice {
	case "fight":
		return GameEvent{Type: Combat, Choices: []PlayerChoice{{Text: "inventory"}, {Text: "fight"}, {Text: "explore"}}}
	case "inventory":
		return GameEvent{Type: ManagingInventory, Choices: []PlayerChoice{{Text: "back"}, {Text: "equip"}}}
	case "explore":
		return GameEvent{Type: Exploring, Choices: []PlayerChoice{{Text: "inventory"}, {Text: "explore"}, {Text: "combat"}}}
	default:
		panic("Bad combat choice")
	}
}

func processInventoryChoice(choice string) GameEvent {
	switch choice {
	case "back":
		if prevState == Combat {
			return GameEvent{Type: Combat, Choices: []PlayerChoice{{Text: "inventory"}, {Text: "fight"}, {Text: "explore"}}}
		} else {
			return GameEvent{Type: Exploring, Choices: []PlayerChoice{{Text: "inventory"}, {Text: "explore"}, {Text: "combat"}}}
		}
	case "equip":
		return GameEvent{Type: ManagingInventory, Choices: []PlayerChoice{{Text: "back"}, {Text: "equip"}}}
	default:
		panic("Bad inventory choice")
	}
}

func updateState(newState GameState) {
	prevState = state
	state = newState
}
