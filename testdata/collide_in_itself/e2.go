package collide_in_itself

import . "github.com/ISKalsi/boomba-the-sapera/models"

var EdgeCaseRequest2 = GameRequest{
	Game: Game{
		ID:      "d75c6192-36ae-42fc-929a-c60b782e721f",
		Timeout: 500,
	},
	Turn: 91,
	Board: Board{
		Height: 7,
		Width:  7,
		Food: []Coord{
			{X: 6, Y: 0},
		},
		Snakes: []Battlesnake{
			{
				ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
				Name:   "Boomba",
				Health: 100,
				Body: []Coord{
					{X: 0, Y: 3},
					{X: 1, Y: 3},
					{X: 2, Y: 3},
					{X: 3, Y: 3},
					{X: 4, Y: 3},
					{X: 4, Y: 4},
					{X: 4, Y: 5},
					{X: 5, Y: 5},
					{X: 5, Y: 6},
					{X: 6, Y: 6},
					{X: 6, Y: 5},
					{X: 6, Y: 4},
					{X: 5, Y: 4},
					{X: 5, Y: 3},
					{X: 5, Y: 2},
					{X: 5, Y: 1},
					{X: 4, Y: 1},
					{X: 3, Y: 1},
					{X: 3, Y: 0},
					{X: 2, Y: 0},
					{X: 1, Y: 0},
					{X: 1, Y: 0},
				},
				Head:   Coord{X: 0, Y: 3},
				Length: 22,
				Shout:  "",
			},
		},
	},
	You: Battlesnake{
		ID:     "3685499c-48cf-4228-8fe0-ae1433968b39",
		Name:   "Boomba",
		Health: 100,
		Body: []Coord{
			{X: 0, Y: 3},
			{X: 1, Y: 3},
			{X: 2, Y: 3},
			{X: 3, Y: 3},
			{X: 4, Y: 3},
			{X: 4, Y: 4},
			{X: 4, Y: 5},
			{X: 5, Y: 5},
			{X: 5, Y: 6},
			{X: 6, Y: 6},
			{X: 6, Y: 5},
			{X: 6, Y: 4},
			{X: 5, Y: 4},
			{X: 5, Y: 3},
			{X: 5, Y: 2},
			{X: 5, Y: 1},
			{X: 4, Y: 1},
			{X: 3, Y: 1},
			{X: 3, Y: 0},
			{X: 2, Y: 0},
			{X: 1, Y: 0},
			{X: 1, Y: 0},
		},
		Head:   Coord{X: 0, Y: 3},
		Length: 22,
		Shout:  "",
	},
}
