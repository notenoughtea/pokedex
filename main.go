package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	internal "github.com/notenoughtea/pokedexcli/internal/pokeAPI"
)

// CLI command structure
type cliCommand struct {
	name        string
	description string
	callback    func(string) error
}

// Commands map
var commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback: func(_ string) error {
			fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Get list of locations and next 20 of them on next use
mapb: Get previous 20 of locations
explore <area>: Explore area
catch <pokemon>: Catch a pokemon
inspect <pokemon>: Inspect a pokemon
pokedex: Display your Pokedex`)
			return nil
		},
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback: func(_ string) error {
			os.Exit(0)
			return nil
		},
	},
	"map": {
		name:        "map",
		description: "Get list of locations and next 20 of them on next use",
		callback: func(_ string) error {
			internal.GetMap(false)
			return nil
		},
	},
	"mapb": {
		name:        "mapb",
		description: "Get previous 20 of locations",
		callback: func(_ string) error {
			internal.GetMap(true)
			return nil
		},
	},
	"explore": {
		name:        "explore",
		description: "Explore a term",
		callback: func(arg string) error {
			if arg == "" {
				return fmt.Errorf("invalid explore command, missing term")
			}
			internal.Explore(arg)
			return nil
		},
	},
	"catch": {
		name:        "catch",
		description: "Catch a pokemon",
		callback: func(arg string) error {
			if arg == "" {
				return fmt.Errorf("invalid catch command, missing term")
			}
			internal.Catch(arg)
			return nil
		},
	},
	"inspect": {
		name:        "inspect",
		description: "inspect cought pokemon",
		callback: func(arg string) error {
			if arg == "" {
				return fmt.Errorf("invalid catch command, missing term")
			}
			internal.Inspect(arg)
			return nil
		},
	},
	"pokedex": {
		name:        "pokedex",
		description: "show your pokedex",
		callback: func(arg string) error {
			internal.Pokedex()
			return nil
		},
	},
}

func main() {
	fmt.Println("Input text:")
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// Print all previously entered lines
		// printLines(lines)

		// Process current line
		if err := processLine(line); err != nil {
			fmt.Println(err)
		}

		// Append current line to history
		lines = append(lines, line)
	}
}

// Prints all lines in the history
// func printLines(lines []string) {
// 	for _, l := range lines {
// 		fmt.Println(l)
// 	}
// }

// Processes the input line and performs appropriate actions
func processLine(line string) error {
	line = strings.TrimSpace(line)
	parts := strings.Fields(line)

	if len(parts) == 0 {
		return nil
	}

	command := parts[0]
	arg := ""
	if len(parts) > 1 {
		arg = strings.Join(parts[1:], " ")
	}

	cmd, exists := commands[command]
	if !exists {
		if len(line) > 0 {
			return fmt.Errorf("unknown command: %s", line)
		}
		return nil
	}

	return cmd.callback(arg)
}
