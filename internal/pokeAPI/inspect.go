package internal

import (
	"fmt"
)

type Stats struct {
	HP             int `json:"hp"`
	Attack         int `json:"attack"`
	Defense        int `json:"defense"`
	SpecialAttack  int `json:"special-attack"`
	SpecialDefense int `json:"special-defense"`
	Speed          int `json:"speed"`
}

type inspectedPokemon struct {
	Name   string   `json:"name"`
	Height int      `json:"height"`
	Weight int      `json:"weight"`
	Stats  Stats    `json:"stats"`
	Types  []string `json:"types"`
}

func Inspect(pokemon string) {
	data := store[pokemon]
	var stats Stats
	for _, stat := range data.Stats {
		switch stat.Stat.Name {
		case "hp":
			stats.HP = stat.BaseStat
		case "attack":
			stats.Attack = stat.BaseStat
		case "defense":
			stats.Defense = stat.BaseStat
		case "special-attack":
			stats.SpecialAttack = stat.BaseStat
		case "special-defense":
			stats.SpecialDefense = stat.BaseStat
		case "speed":
			stats.Speed = stat.BaseStat
		}
	}

	var types []string
	for _, type_ := range data.Types {
		types = append(types, type_.Type.Name)
	}

	targetedPokemon := inspectedPokemon{
		Name:   data.Name,
		Height: data.Height,
		Weight: data.Weight,
		Stats:  stats,
		Types:  types,
	}

	fmt.Printf("Inspecting %s\n", pokemon)

	fmt.Printf(`Name: %s
Height: %d
Weight: %d
Stats:
  -HP: %d
  -Attack: %d
  -Defense: %d
  -Special Attack: %d
  -Special Defense: %d
  -Speed: %d
Types:
`,
		targetedPokemon.Name,
		targetedPokemon.Height,
		targetedPokemon.Weight,
		targetedPokemon.Stats.HP,
		targetedPokemon.Stats.Attack,
		targetedPokemon.Stats.Defense,
		targetedPokemon.Stats.SpecialAttack,
		targetedPokemon.Stats.SpecialDefense,
		targetedPokemon.Stats.Speed,
	)
	for _, item := range types {
		fmt.Printf("-%s\n", item)
	}

}
