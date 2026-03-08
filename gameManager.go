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

type Choice interface {
}

type PlayerChoice struct {
	DisplayText string
	Aliases     []string
	raw         string
	payload     interface{}
}

var ExploreEvent = GameEvent{Type: Exploring}
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

		starterEvent := ExploreEvent
		starterEvent.Description = getDirectionDescriptions()
		starterEvent.Choices = generateExplorationChoices()

		events <- starterEvent

		for choice := range input {
			switch state {
			case Exploring:
				newEvent = processExplorationChoice(choice.raw)
			case Combat:
				newEvent = processCombatChoice(choice.raw)
			case ManagingInventory:
				newEvent = processInventoryChoice(choice.raw)
			}

			fp.Illustrate(playerPosition, fp.Point{}, level.Grid)

			updateState(newEvent.Type)

			events <- newEvent
		}

	}()

	return events
}

func generateExplorationChoices() []PlayerChoice {
	exits := fp.GetExitDirections(playerPosition.X, playerPosition.Y, level.Grid)
	choices := make([]PlayerChoice, 0)
	for _, direction := range exits {
		choices = append(choices, DirectionalChoices[direction])
	}

	return append(choices, manageInventoryChoice)
}

func getDirectionDescriptions() string {
	directions := fp.GetExitDirections(playerPosition.X, playerPosition.Y, level.Grid)
	directionText := make([]string, 0)
	for _, direction := range directions {
		switch direction {
		case fp.North:
			directionText = append(directionText, "North")
		case fp.South:
			directionText = append(directionText, "South")
		case fp.East:
			directionText = append(directionText, "East")
		case fp.West:
			directionText = append(directionText, "West")
		}
	}

	switch len(directionText) {
	case 1:
		return directionText[0]
	case 2:
		return fmt.Sprintf("%s and %s", directionText[0], directionText[1])
	case 3:
		return fmt.Sprintf("%s, %s, and %s", directionText[0], directionText[1], directionText[2])
	case 4:
		return fmt.Sprintf("%s, %s, %s, and %s", directionText[0], directionText[1], directionText[2], directionText[3])
	default:
		panic("Ahhhhhh there are either zero or more than four exits, what is happening")
	}

}

func processExplorationChoice(choice string) GameEvent {
	switch choice {
	case north:
		playerPosition.Y += 1
		event := ExploreEvent
		event.Description = getDirectionDescriptions()
		event.Choices = generateExplorationChoices()
		return event
	case south:
		playerPosition.Y -= 1
		event := ExploreEvent
		event.Description = getDirectionDescriptions()
		event.Choices = generateExplorationChoices()
		return event
	case east:
		playerPosition.X += 1
		event := ExploreEvent
		event.Description = getDirectionDescriptions()
		event.Choices = generateExplorationChoices()
		return event
	case west:
		playerPosition.X -= 1
		event := ExploreEvent
		event.Description = getDirectionDescriptions()
		event.Choices = generateExplorationChoices()
		return event
	case inventory:
		return ManageInventoryEvent
	default:
		panic("Bad exploration choice")
	}
}

func processCombatChoice(choice string) GameEvent {
	switch choice {
	case inventory:
		return ManageInventoryEvent
	default:
		panic("Bad combat choice")
	}
}

func processInventoryChoice(choice string) GameEvent {
	switch choice {
	case back:
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
