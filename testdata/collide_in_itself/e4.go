package collide_in_itself

import . "github.com/ISKalsi/boomba-the-sapera/models"

var EdgeCaseRequest4 = GameRequest{
	Game: Game{
		ID:      "4e6026ec-b688-46a8-8842-473b50acf58c",
		Timeout: 500,
	},
	Turn: 115,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 4, Y: 0},
		},
		Snakes: []Battlesnake{
			{
				ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
				Name:   "Boomba",
				Health: 94,
				Body: []Coord{
					{X: 5, Y: 8},
					{X: 5, Y: 7},
					{X: 5, Y: 6},
					{X: 6, Y: 6},
					{X: 6, Y: 5},
					{X: 5, Y: 5},
					{X: 5, Y: 4},
					{X: 6, Y: 4},
					{X: 7, Y: 4},
					{X: 7, Y: 5},
					{X: 7, Y: 6},
					{X: 7, Y: 7},
					{X: 8, Y: 7},
					{X: 9, Y: 7},
					{X: 9, Y: 6},
					{X: 9, Y: 5},
					{X: 9, Y: 4},
				},
				Head:   Coord{X: 5, Y: 8},
				Length: 17,
				Shout:  "",
			},
			{
				ID:     "1355499c-48cf-4228-8fe0-ae1433968b14",
				Name:   "Ava",
				Health: 79,
				Body: []Coord{
					{X: 4, Y: 1},
					{X: 4, Y: 2},
					{X: 4, Y: 3},
					{X: 3, Y: 3},
					{X: 3, Y: 4},
					{X: 3, Y: 5},
					{X: 3, Y: 6},
					{X: 2, Y: 6},
					{X: 1, Y: 6},
					{X: 1, Y: 5},
					{X: 1, Y: 4},
					{X: 1, Y: 3},
					{X: 0, Y: 3},
				},
				Head:   Coord{X: 4, Y: 1},
				Length: 13,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
		Name:   "Boomba",
		Health: 94,
		Body: []Coord{
			{X: 5, Y: 8},
			{X: 5, Y: 7},
			{X: 5, Y: 6},
			{X: 6, Y: 6},
			{X: 6, Y: 5},
			{X: 5, Y: 5},
			{X: 5, Y: 4},
			{X: 6, Y: 4},
			{X: 7, Y: 4},
			{X: 7, Y: 5},
			{X: 7, Y: 6},
			{X: 7, Y: 7},
			{X: 8, Y: 7},
			{X: 9, Y: 7},
			{X: 9, Y: 6},
			{X: 9, Y: 5},
			{X: 9, Y: 4},
		},
		Head:   Coord{X: 5, Y: 8},
		Length: 17,
		Shout:  "",
	},
}
