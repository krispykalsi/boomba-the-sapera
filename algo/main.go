package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/algo/coord"
	"github.com/ISKalsi/boomba-the-sapera/models"
)

type Algorithm struct {
	board       models.Board
	start       models.Coord
	destination models.Coord
	solvedPath  []models.Coord
	isSolving   bool
	head        models.Coord
}

type PathSolver interface {
	NextMove(gr *models.GameRequest) string
}

func Init(b models.Board) *Algorithm {
	s := b.Snakes[0]
	return &Algorithm{
		board:       b,
		start:       s.Head,
		head:        s.Head,
		destination: b.Food[0],
		solvedPath:  make([]models.Coord, 0),
		isSolving:   false,
	}
}

func (a *Algorithm) reset(b models.Board) {
	s := b.Snakes[0]
	a.board = b
	a.start = s.Head
	a.head = s.Head
	a.destination = b.Food[0]
	a.solvedPath = a.solvedPath[:0]
	a.isSolving = false
}

func (a *Algorithm) getNextDirection() string {
	next := a.solvedPath[0]
	a.solvedPath = a.solvedPath[1:]
	dir := coord.Diff(&next, &a.head)
	a.head = next
	return parseMoveDirectionToString(directionToIndex[dir])
}

func (a *Algorithm) NextMove(gr *models.GameRequest) string {
	if len(a.solvedPath) != 0 {
		return a.getNextDirection()
	} else if a.isSolving {
		return getRandomMove("solving...: ")
	} else {
		a.reset(gr.Board)
		if a.aStarSearch() {
			return a.getNextDirection()
		} else {
			return getRandomMove("random move: ")
		}
	}
}