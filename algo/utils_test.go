package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/models"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestParseMoveIndexToString(t *testing.T) {
	tests := []struct {
		expectedString string
		moveIndex      int
	}{
		{"up", 0},
		{"down", 1},
		{"left", 2},
		{"right", 3},
	}

	for _, test := range tests {
		t.Run(strings.ToUpper(test.expectedString), func(t *testing.T) {
			m := parseMoveDirectionToString(test.moveIndex)
			assert.Equal(t, test.expectedString, m)
		})
	}
}

func TestPanicFromParseMoveIndexToString(t *testing.T) {
	assert.Panics(t, func() {
		_ = parseMoveDirectionToString(4)
	})
}

func TestDirectionAndIndexMapEquality(t *testing.T) {
	for dir, idx := range directionToIndex {
		assert.Equal(t, indexToDirection[idx], dir)
	}
}

func TestInsert(t *testing.T) {
	a := models.Coord{X: 0, Y: 0}
	b := models.Coord{X: 0, Y: 1}
	c := models.Coord{X: 1, Y: 0}
	d := models.Coord{X: 1, Y: 1}

	tests := []struct {
		index int
	}{
		{0},
		{1},
		{2},
		{3},
	}

	for _, test := range tests {
		t.Run("index"+strconv.Itoa(test.index), func(t *testing.T) {
			cells := []models.Coord{a, b, c}
			cells = insert(cells, test.index, d)

			assert.Equal(t, d, cells[test.index])
		})
	}
}

func TestReverse(t *testing.T) {
	a := models.Coord{X: 0, Y: 0}
	b := models.Coord{X: 0, Y: 1}
	c := models.Coord{X: 1, Y: 0}
	d := models.Coord{X: 1, Y: 1}

	slice := []models.Coord{a, b, c, d}
	reversedSlice := []models.Coord{d, c, b, a}

	reverse(slice)

	assert.Equal(t, reversedSlice, slice)
}
