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
	DisplayText string
	Aliases     []string
	raw         string
	payload     interface{}
}

var ExploreEvent = GameEvent{Type: Exploring, Choices: []PlayerChoice{{DisplayText: "[I]nventory", raw: "inventory", Aliases: []string{"i", "inventory"}}}}
var CombatEvent = GameEvent{Type: Combat, Choices: []PlayerChoice{{DisplayText: "[I]nventory", raw: "inventory", Aliases: []string{"i", "inventory"}}}}
var ManageInventoryEvent = GameEvent{Type: ManagingInventory, Choices: []PlayerChoice{{DisplayText: "[B]ack", raw: "back", Aliases: []string{"b", "back"}}}}

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
		availableChoices := ExploreEvent.Choices

		for choice := range input {
			switch state {
			case Exploring:
				newEvent = processExplorationChoice(choice.DisplayText, availableChoices)
			case Combat:
				newEvent = processCombatChoice(choice.DisplayText, availableChoices)
			case ManagingInventory:
				newEvent = processInventoryChoice(choice.DisplayText, availableChoices)
			}

			var hold string
			fmt.Scanln(&hold)

			updateState(newEvent.Type)
			availableChoices = newEvent.Choices

			events <- newEvent
		}

	}()

	return events
}

func processExplorationChoice(choice string, availableChoices []PlayerChoice) GameEvent {
	switch choice {
	case "inventory":
		return ManageInventoryEvent
	default:
		panic("Bad exploration choice")
	}
}

func processCombatChoice(choice string, availableChoices []PlayerChoice) GameEvent {
	switch choice {
	case "inventory":
		return ManageInventoryEvent
	default:
		panic("Bad combat choice")
	}
}

func processInventoryChoice(choice string, availableChoices []PlayerChoice) GameEvent {
	switch choice {
	case "back":
		if prevState == Combat {
			return CombatEvent
		} else {
			return ExploreEvent
		}
	default:
		panic("Bad inventory choice")
	}
}

func updateState(newState GameState) {
	prevState = state
	state = newState
}
