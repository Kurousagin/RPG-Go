package provider

import (
	
)

const (
	goblin = "Goblin"
	orc    = "Orc"
	trol   = "Trol"
	dragon = "Dragon"
)

var monsters = map[string]map[string]int{
	goblin: {"Dano": 5},
	orc:    {"Dano": 7},
	trol:   {"Dano": 10},
	dragon: {"Dano": 15},
}
