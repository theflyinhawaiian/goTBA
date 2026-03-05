package main

import (
	"fmt"
	"tba/entities"
	"tba/floorplan"
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

var ExploreEvent = GameEvent{Type: Exploring, Choices: []PlayerChoice{{Text: "[I]nventory"}, {Text: "[E]xplore"}, {Text: "[C]ombat"}}}
var CombatEvent = GameEvent{Type: Combat, Choices: []PlayerChoice{{Text: "[I]nventory"}, {Text: "[E]xplore"}, {Text: "[F]ight"}}}
var ManageInventoryEvent = GameEvent{Type: ManagingInventory, Choices: []PlayerChoice{{Text: "[B]ack"}, {Text: "[E]quip"}}}

var level fp.Map
var state GameState
var prevState GameState
var player entities.Player
var playerPosition floorplan.Point

func Start(input chan PlayerChoice) <-chan GameEvent {
	level = fp.GenerateLevel()
	player = entities.CreatePlayer()
	playerPosition = level.Start
	updateState(Exploring)

	events := make(chan GameEvent)
	go func() {
		defer close(events)
		var newEvent GameEvent

		events <- ExploreEvent

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
		return ExploreEvent
	case "combat":
		return CombatEvent
	case "inventory":
		return ManageInventoryEvent
	default:
		panic("Bad exploration choice")
	}
}

func processCombatChoice(choice string) GameEvent {
	switch choice {
	case "fight":
		return CombatEvent
	case "inventory":
		return ManageInventoryEvent
	case "explore":
		return ExploreEvent
	default:
		panic("Bad combat choice")
	}
}

func processInventoryChoice(choice string) GameEvent {
	switch choice {
	case "back":
		if prevState == Combat {
			return CombatEvent
		} else {
			return ExploreEvent
		}
	case "equip":
		return ManageInventoryEvent
	default:
		panic("Bad inventory choice")
	}
}

func updateState(newState GameState) {
	prevState = state
	state = newState
}
