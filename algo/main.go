package algo

import (
	"github.com/ISKalsi/boomba-the-sapera/algo/cell"
	"github.com/ISKalsi/boomba-the-sapera/algo/grid"
	"github.com/ISKalsi/boomba-the-sapera/models"
	"math"
	"sort"
)

type PathSolver interface {
	NextMove(gr *models.GameRequest) string
}

type Algorithm struct {
	board               models.Board
	start               models.Coord
	destination         models.Coord
	solvedPath          []models.Coord
	head                models.Coord
	health              float64
	dontBlockTailOrHead bool
	dontConsiderCost    bool
	ignoreUnsureBlocks  bool
	headCollisions      possibleHeadCollisions
}

func Init(b models.Board, s models.Battlesnake) *Algorithm {
	return &Algorithm{
		board:          b,
		start:          s.Head,
		head:           s.Head,
		destination:    s.Head.FindNearestCoordsFrom(b.Food)[0],
		health:         float64(s.Health),
		solvedPath:     make([]models.Coord, 0),
		headCollisions: possibleHeadCollisions{coords: []models.Coord{}},
	}
}

func (a *Algorithm) clearSolvedPath() {
	a.solvedPath = a.solvedPath[:0]
}

func (a *Algorithm) SetNewStart(c models.Coord) {
	a.start = c
}

func (a *Algorithm) SetNewDestination(c models.Coord) {
	a.destination = c
}

func (a *Algorithm) findPossibleLosingHeadCollisions(ourSnake models.Battlesnake) {
	var dangerCoords []models.Coord
	for _, opponent := range a.board.Snakes {
		if opponent.Length >= ourSnake.Length {
			h := ourSnake.Head.CalculateHeuristics(opponent.Head)
			if h == 2 {
				for dir := range directionToIndex {
					c := opponent.Head.Sum(dir)
					if !c.IsOutside(a.board.Width, a.board.Height) {
						dangerCoords = append(dangerCoords, c)
					}
				}
			}
		}
	}

	a.headCollisions.coords = dangerCoords
}

func (a *Algorithm) sortFoodByHeuristicFrom(head models.Coord) {
	sort.Slice(a.board.Food, func(i, j int) bool {
		h1 := a.board.Food[i].CalculateHeuristics(head)
		h2 := a.board.Food[j].CalculateHeuristics(head)
		return h1 < h2
	})
}

func (a *Algorithm) findNearestPlausibleFood() (bool, models.Coord) {
	a.sortFoodByHeuristicFrom(a.head)

	heads := make([]models.Coord, len(a.board.Snakes))
	for i, snake := range a.board.Snakes {
		heads[i] = snake.Head
	}

	for _, foodCoord := range a.board.Food {
		nearestHeads := foodCoord.FindNearestCoordsFrom(heads)
		if len(nearestHeads) == 1 && nearestHeads[0] == a.head {
			return true, foodCoord
		}
	}

	return false, a.board.Food[0]
}

func (a *Algorithm) reset(b models.Board, s models.Battlesnake) {
	a.board = b
	a.start = s.Head
	a.head = s.Head
	a.health = float64(s.Health)
	a.clearSolvedPath()
}

func (a *Algorithm) getDirection(next models.Coord) string {
	dir := next.Diff(a.head)
	a.head = next
	return parseMoveDirectionToString(directionToIndex[dir])
}

func (a *Algorithm) initGrid() grid.Grid {
	obstacles := make([]grid.ObstacleProvider, len(a.board.Snakes))
	for i, snake := range a.board.Snakes {
		obstacles[i] = snake
	}

	maybeObstacles := make([]grid.PotentialObstacleProvider, 1)
	maybeObstacles[0] = a.headCollisions

	g := grid.WithObstacles(a.board.Width, a.board.Height, obstacles, maybeObstacles)

	for _, hazardCoord := range a.board.Hazards {
		g[hazardCoord].Weight = cell.WeightHazard
	}

	return g
}

func (a *Algorithm) getTrappedScore(g grid.Grid, gr *models.GameRequest) int {
	trappedScore := 0
	for dir := range directionToIndex {
		test := gr.You.Head.Sum(dir)

		if isOwnBody := gr.You.Body[1] == test; isOwnBody {
			continue
		}

		if test.IsOutside(a.board.Width, a.board.Height) {
			trappedScore += 1
		} else if g[test].Weight == cell.WeightHazard {
			trappedScore += 1
		} else if g[test].IsBlocked {
			trappedScore += 1
		}
	}
	return trappedScore
}

func (a *Algorithm) NextMove(gr *models.GameRequest) string {
	a.reset(gr.Board, gr.You)
	a.findPossibleLosingHeadCollisions(gr.You)
	g := a.initGrid()

	pathFound := false
	pathCost := 100.0

	nearestFoodFound, foodCoord := a.findNearestPlausibleFood()
	a.destination = foodCoord

	if nearestFoodFound {
		pathFound, pathCost = a.aStarSearch()
	}

	if pathFound && (gr.You.Length < 15 || float64(gr.You.Health) <= pathCost) {
		shortestPathNextCoord := a.solvedPath[0]
		trappedScore := a.getTrappedScore(g, gr)
		virtualSnake := g.MoveVirtualSnakeAlongPath(gr.You.Body, a.solvedPath)

		ourSnakeIndex := 0
		originalSnakeBody := gr.You.Body[:]
		originalSnakeHealth := a.health

		for i := range a.board.Snakes {
			if a.board.Snakes[i].ID == gr.You.ID {
				ourSnakeIndex = i
				a.board.Snakes[i].Body = virtualSnake
				a.board.Snakes[i].Head = virtualSnake[0]
				if g[foodCoord].Weight == cell.WeightHazard {
					a.health = 100 - cell.WeightHazard
				} else {
					a.health = 100
				}
				break
			}
		}

		a.SetNewStart(virtualSnake[0])
		a.SetNewDestination(virtualSnake[len(virtualSnake)-1])
		a.dontBlockTailOrHead = true

		if pathFound, pathCost = a.aStarSearch(); pathFound {
			if a.health-pathCost < 35 && originalSnakeHealth <= 30 || trappedScore > 1 {
				return a.getDirection(shortestPathNextCoord)
			} else if len(a.solvedPath) != 1 || !foodCoord.IsAtEdge(a.board.Width, a.board.Height) {
				return a.getDirection(shortestPathNextCoord)
			}
		}

		a.board.Snakes[ourSnakeIndex].Body = originalSnakeBody
		a.board.Snakes[ourSnakeIndex].Head = originalSnakeBody[0]
		a.health = originalSnakeHealth
	}

	g = a.initGrid()
	var nearestSnake models.Battlesnake
	nearestSnakeIndex := -1
	minH := math.Inf(1)
	for i, snake := range a.board.Snakes {
		if snake.ID != gr.You.ID {
			if h := gr.You.Head.CalculateHeuristics(snake.Head); minH > h {
				minH = h
				nearestSnake = snake
				nearestSnakeIndex = i
			}
		}
	}

	bodyPartsInHazard := 0
	isHeadInHazard := g[gr.You.Head].Weight == cell.WeightHazard
	for _, c := range gr.You.Body {
		if g[c].Weight == cell.WeightHazard {
			bodyPartsInHazard += 1
		}
	}
	bodyInHazardToSafeRatio := float32(bodyPartsInHazard) / float32(gr.You.Length)

	pathFoundIsTooCostly := false
	if len(a.board.Snakes) == 1 {
		if found, direction := a.findLongestPathToTail(&gr.You); found {
			return direction
		}
	} else if bodyInHazardToSafeRatio < 0.2 || isHeadInHazard {
		bigSnakesAround := false
		collisionPosibility := false

		a.SetNewStart(gr.You.Head)
		a.SetNewDestination(gr.You.Body[len(gr.You.Body)-1])
		a.dontBlockTailOrHead = true

		for _, snake := range a.board.Snakes {
			if snake.ID != gr.You.ID && snake.Length >= gr.You.Length {
				if h := snake.Head.CalculateHeuristics(gr.You.Head); h < 6 {
					if pathFound, pathCost = a.aStarSearch(); pathFound {
						pathLen := len(a.solvedPath)
						if pathCost >= 45 && pathLen < 5 {
							pathFoundIsTooCostly = true
						} else {
							tailIndex := gr.You.Length - 1
							justHadFood := gr.You.Body[tailIndex] == gr.You.Body[tailIndex-1]
							if pathLen != 1 || !justHadFood {
								return a.getDirection(a.solvedPath[0])
							}
						}
					}
					collisionPosibility = true
				}
				bigSnakesAround = true
			}
		}

		if !bigSnakesAround {
			a.SetNewStart(gr.You.Head)
			a.SetNewDestination(nearestSnake.Head)
			if pathFound, pathCost = a.aStarSearch(); pathFound {
				shortestPathNextCoord := a.solvedPath[0]
				virtualSnake := g.MoveVirtualSnakeAlongPath(gr.You.Body, a.solvedPath[:1])

				ourSnakeIndex := 0
				originalSnakeBody := gr.You.Body[:]
				originalSnakeHealth := a.health

				for i := range a.board.Snakes {
					if a.board.Snakes[i].ID == gr.You.ID {
						ourSnakeIndex = i
						a.board.Snakes[i].Body = virtualSnake
						a.board.Snakes[i].Head = virtualSnake[0]
						a.health -= pathCost
						break
					}
				}

				willBeOkay := true
				a.SetNewStart(virtualSnake[0])

				originalNearestSnakeBody := nearestSnake.Body[:]
				a.health -= 1

				for dir := range directionToIndex {
					test := nearestSnake.Head.Sum(dir)

					isOwnBody := nearestSnake.Body[1] == test
					if isOwnBody || test.IsOutside(a.board.Width, a.board.Height) || !g[test].IsOk() {
						continue
					}

					testVirtualSnake := g.MoveVirtualSnakeAlongPath(nearestSnake.Body, []models.Coord{test})
					a.board.Snakes[nearestSnakeIndex].Body = testVirtualSnake
					a.board.Snakes[nearestSnakeIndex].Head = testVirtualSnake[0]

					a.SetNewDestination(test)
					a.dontBlockTailOrHead = true
					pathFound, pathCost = a.aStarSearch()
					if !pathFound || g[a.solvedPath[0]].Weight == cell.WeightHazard {
						willBeOkay = false
						break
					}
				}

				a.board.Snakes[nearestSnakeIndex].Body = originalNearestSnakeBody
				a.board.Snakes[nearestSnakeIndex].Head = originalNearestSnakeBody[0]
				a.health += 1

				if willBeOkay {
					return a.getDirection(shortestPathNextCoord)
				}

				a.board.Snakes[ourSnakeIndex].Body = originalSnakeBody
				a.board.Snakes[ourSnakeIndex].Head = originalSnakeBody[0]
				a.health = originalSnakeHealth
			}
		} else {
			a.SetNewStart(gr.You.Head)
			a.SetNewDestination(nearestSnake.Body[nearestSnake.Length-1])
			if pathFound, pathCost = a.aStarSearch(); pathFound {
				tailIndex := nearestSnake.Length - 1
				justHadFood := nearestSnake.Body[tailIndex] == nearestSnake.Body[tailIndex-1]
				if len(a.solvedPath) != 1 || !justHadFood {
					return a.getDirection(a.solvedPath[0])
				}
			}
		}

		if !collisionPosibility {
			if found, direction := a.findLongestPathToTail(&gr.You); found {
				return direction
			}
		}
	} else if nearestSnakeIndex != -1 {
		a.SetNewStart(gr.You.Head)
		a.SetNewDestination(nearestSnake.Body[nearestSnake.Length-1])
		if pathFound, pathCost = a.longestPath(); pathFound {
			tailIndex := nearestSnake.Length - 1
			justHadFood := nearestSnake.Body[tailIndex] == nearestSnake.Body[tailIndex-1]
			if len(a.solvedPath) != 1 || !justHadFood {
				return a.getDirection(a.solvedPath[0])
			}
		}
	}

	g = a.initGrid()
	var avoidCoord models.Coord
	notFoundAvoidCoord := true

	if pathFoundIsTooCostly {
		for _, snake := range a.board.Snakes {
			if snake.Length >= gr.You.Length {
				avoidCoord = snake.Head
				notFoundAvoidCoord = false
				break
			}
		}
	} else {
		avoidCoord = foodCoord
	}

	minF := math.Inf(1)
	var maxDir models.Coord

	if !notFoundAvoidCoord {
		for dir := range directionToIndex {
			test := gr.You.Head.Sum(dir)

			isOwnBody := gr.You.Body[1] == test
			if isOwnBody || test.IsOutside(a.board.Width, a.board.Height) {
				continue
			}

			if !g[test].IsOk() {
				continue
			}

			H := g[test].CalculateHeuristics(avoidCoord)
			G := g[test].Weight
			F := G - H
			if F <= minF {
				a.SetNewStart(test)
				a.SetNewDestination(gr.You.Body[len(gr.You.Body)-1])
				a.dontBlockTailOrHead = true
				a.ignoreUnsureBlocks = true
				if pathFound, _ = a.aStarSearch(); pathFound {
					tailIndex := gr.You.Length - 1
					justHadFood := gr.You.Body[tailIndex] == gr.You.Body[tailIndex-1]
					if len(a.solvedPath) != 1 || !justHadFood {
						minF = F
						maxDir = dir
					}
				}
			}
		}
	}

	var goTowardsCoord models.Coord
	if nearestSnakeIndex != -1 {
		goTowardsCoord = nearestSnake.Body[nearestSnake.Length-1]
	} else {
		goTowardsCoord = gr.You.Body[gr.You.Length-1]
	}

	if minF == math.Inf(1) {
		cost := math.Inf(1)
		for dir := range directionToIndex {
			test := gr.You.Head.Sum(dir)

			isOwnBody := gr.You.Body[1] == test
			if isOwnBody || test.IsOutside(a.board.Width, a.board.Height) {
				continue
			}

			if g[test].IsBlocked {
				continue
			}

			H := g[test].CalculateHeuristics(goTowardsCoord)
			G := g[test].Weight
			F := G + H
			if F <= minF {
				a.SetNewStart(test)
				a.SetNewDestination(gr.You.Body[len(gr.You.Body)-1])
				a.dontBlockTailOrHead = true
				a.dontConsiderCost = true
				if pathFound, pathCost = a.aStarSearch(); pathFound {
					if pathCost < cost {
						cost = pathCost
						minF = F
						maxDir = dir
					}
				} else if minF == math.Inf(1) {
					minF = F
					maxDir = dir
				}
			}
		}
	}

	return parseMoveDirectionToString(directionToIndex[maxDir])
}

func (a *Algorithm) findLongestPathToTail(you *models.Battlesnake) (bool, string) {
	a.SetNewStart(you.Head)
	a.SetNewDestination(you.Body[len(you.Body)-1])
	a.dontBlockTailOrHead = true
	if pathFound, _ := a.longestPath(); pathFound {
		tailIndex := you.Length - 1
		justHadFood := you.Body[tailIndex] == you.Body[tailIndex-1]
		if len(a.solvedPath) != 1 || !justHadFood {
			return true, a.getDirection(a.solvedPath[0])
		}
	}
	return false, ""
}
