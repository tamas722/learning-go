package basics

import (
	"errors"
	"fmt"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR STRUCTS HERE
type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

// newGame returns a new game struct.
func newGame(id int, name string, price int, genre string) game {
	// INSERT YOUR CODE HERE
	return game{
		item: item{
			id:    id,
			name:  name,
			price: price,
		},
		genre: genre,
	}
}

// String stringifies an item.
func (i item) String() string {
	// INSERT YOUR CODE HERE
	return fmt.Sprintf("%d: %s costs %d", i.id, i.name, i.price)
}

// String stringifies a game.
func (g game) String() string {
	// INSERT YOUR CODE HERE
	return fmt.Sprintf("Game %s of genre %s", g.item.String(), g.genre)
}

// newGameList creates a game store.
func newGameList() []game {
	// INSERT YOUR CODE HERE
	return []game{
		newGame(2, "x-com 2", 30, "strategy"),
		newGame(3, "minecraft", 20, "sandbox"),
		newGame(4, "warcraft", 40, "strategy"),
	}
}

// queryById returns the game in the specified store with the given id or returns a "no such game" error.
func queryById(games []game, id int) (game, error) {
	// INSERT YOUR CODE HERE
	for _, g := range games {
		if g.id == id {
			return g, nil
		}
	}
	return game{}, errors.New("no such game")
}

// listNameByPrice returns the name of the game(s) with price equal or smaller than a given price.
func listNameByPrice(games []game, price int) []string {
	// INSERT YOUR CODE HERE
	var names []string
	for _, g := range games {
		if g.price <= price {
			names = append(names, g.name)
		}
	}
	return names
}
