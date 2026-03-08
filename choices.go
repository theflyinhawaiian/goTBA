package main

import "tba/entities"

const (
	north     = "north"
	south     = "south"
	east      = "east"
	west      = "west"
	inventory = "inventory"
	attack    = "attack"
	cast      = "cast"
	heal      = "heal"
	equip     = "equip"
	back      = "back"
)

var northChoice = PlayerChoice{DisplayText: "Go [N]orth", Aliases: []string{"n", "north"}, raw: north}
var southChoice = PlayerChoice{DisplayText: "Go [S]outh", Aliases: []string{"s", "south"}, raw: south}
var eastChoice = PlayerChoice{DisplayText: "Go [E]ast", Aliases: []string{"e", "east"}, raw: east}
var westChoice = PlayerChoice{DisplayText: "Go [W]est", Aliases: []string{"w", "west"}, raw: west}

var DirectionalChoices = []PlayerChoice{northChoice, southChoice, eastChoice, westChoice}

var manageInventoryChoice = PlayerChoice{DisplayText: "[I]nventory", raw: "inventory", Aliases: []string{"i", inventory}}

var MeleeChoice = PlayerChoice{DisplayText: "[A]ttack", Aliases: []string{"a", "attack"}, raw: attack}
var MagicChoice = PlayerChoice{DisplayText: "[C]ast", Aliases: []string{"c", "cast"}, raw: cast}
var HealChoice = PlayerChoice{DisplayText: "[H]eal", Aliases: []string{"h", "heal"}, raw: heal}

func generateEquipChoice(item entities.Item) PlayerChoice {
	return PlayerChoice{DisplayText: item.Name, Aliases: []string{item.Name}, raw: equip, payload: item.Id}
}
