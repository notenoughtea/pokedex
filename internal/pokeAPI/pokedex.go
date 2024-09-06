package internal

import (
	"fmt"
)

func Pokedex() {
	if len(store) == 0 {
		fmt.Println("No pokemon in Pokedex")
		return
	}
	fmt.Println("Your Pokedex:")
	for _, k := range store {
		fmt.Printf("- %v\n", k.Name)
	}
}
