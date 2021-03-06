package hazard_related

import (
	. "github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/ISKalsi/boomba-the-sapera/testdata"
)

var EdgeCaseRequest4 = GameRequest{
	Game: Game{
		ID:      "30f24d94-9ebe-4798-a75b-dd450af6d6ed",
		Timeout: 500,
	},
	Turn: 58,
	Board: Board{
		Height: 11,
		Width:  11,
		Food: []Coord{
			{X: 2, Y: 5},
		},
		Snakes: []Battlesnake{
			{
				ID:     "gs_dCfHTYqrx6dvHBwfphRkWyKF",
				Name:   "Boomba",
				Health: 94,
				Body: []Coord{
					{X: 1, Y: 3},
					{X: 2, Y: 3},
					{X: 3, Y: 3},
					{X: 4, Y: 3},
					{X: 5, Y: 3},
					{X: 6, Y: 3},
					{X: 6, Y: 4},
					{X: 7, Y: 4},
				},
				Head:   Coord{X: 1, Y: 3},
				Length: 8,
				Shout:  "",
			},
			{
				ID:     "gs_HkjRFJFBvWRW6RyShgdXR4kC",
				Name:   "snake30series",
				Health: 99,
				Body: []Coord{
					{X: 2, Y: 6},
					{X: 2, Y: 5},
					{X: 3, Y: 5},
					{X: 4, Y: 5},
					{X: 5, Y: 5},
					{X: 6, Y: 5},
					{X: 6, Y: 6},
					{X: 6, Y: 7},
					{X: 6, Y: 8},
					{X: 6, Y: 9},
					{X: 5, Y: 9},
				},
				Head:   Coord{X: 2, Y: 6},
				Length: 11,
				Shout:  "",
			},
		},
		Hazards: append(
			testdata.CreateHazardCoordsFromColumns(11, 0),
			testdata.CreateHazardCoordsFromRows(11, 10)...,
		),
	},
	You: Battlesnake{
		ID:     "gs_dCfHTYqrx6dvHBwfphRkWyKF",
		Name:   "Boomba",
		Health: 94,
		Body: []Coord{
			{X: 1, Y: 3},
			{X: 2, Y: 3},
			{X: 3, Y: 3},
			{X: 4, Y: 3},
			{X: 5, Y: 3},
			{X: 6, Y: 3},
			{X: 6, Y: 4},
			{X: 7, Y: 4},
		},
		Head:   Coord{X: 1, Y: 3},
		Length: 8,
		Shout:  "",
	},
}
