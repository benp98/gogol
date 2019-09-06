// Package gogol is an implementation of Conway's Game of Life
package gogol

// State is an instance of the Game of Life
type State struct {
	width           int
	height          int
	world           [][]bool
	previousWorld   [][]bool
	neighbourRadius int
}

// NewState creates a new Game of Life State
func NewState(width, height, neighbourRadius int) *State {
	state := new(State)
	state.width = width
	state.height = height
	state.neighbourRadius = neighbourRadius

	state.world = makeWorld(width, height)

	return state
}

// GetDimensions returns the dimensions of the world
func (state *State) GetDimensions() (int, int) {
	return state.width, state.height
}

// SetCell sets the state of the specified cell
func (state *State) SetCell(x, y int, value bool) {
	nx, ny := state.normalizeCoordinates(x, y)
	state.world[ny][nx] = value
}

// GetCell returns the state of the specified cell
func (state *State) GetCell(x, y int) bool {
	nx, ny := state.normalizeCoordinates(x, y)
	return state.world[ny][nx]
}

// NextGeneration calculates the next generation of the game
func (state *State) NextGeneration() {
	state.previousWorld = state.world
	state.world = makeWorld(state.width, state.height)

	for y := 0; y < state.height; y++ {
		for x := 0; x < state.width; x++ {
			aliveNeighbours := state.countAliveNeighbours(x, y)
			cellState := state.getPreviousGenerationCell(x, y)

			switch true {
			case (!cellState && aliveNeighbours == 3):
				state.SetCell(x, y, true)
			case cellState && aliveNeighbours < 2:
				state.SetCell(x, y, false)
			case cellState && (aliveNeighbours == 2 || aliveNeighbours == 3):
				state.SetCell(x, y, true)
			case cellState && aliveNeighbours > 3:
				state.SetCell(x, y, false)
			}
		}
	}
}

func (state *State) normalizeCoordinates(x, y int) (int, int) {
	nx := x
	for nx < 0 {
		nx += state.width
	}
	for nx >= state.width {
		nx -= state.width
	}

	ny := y
	for ny < 0 {
		ny += state.height
	}
	for ny >= state.height {
		ny -= state.height
	}

	return nx, ny
}

func (state *State) getPreviousGenerationCell(x, y int) bool {
	nx, ny := state.normalizeCoordinates(x, y)
	return state.previousWorld[ny][nx]
}

func (state *State) countAliveNeighbours(x, y int) int {
	aliveNeighbours := 0

	for i := -state.neighbourRadius; i <= state.neighbourRadius; i++ {
		for j := -state.neighbourRadius; j <= state.neighbourRadius; j++ {
			if state.getPreviousGenerationCell(x+i, y+j) {
				aliveNeighbours++
			}
		}
	}

	return aliveNeighbours
}

func makeWorld(width, height int) [][]bool {
	world := make([][]bool, height)
	for i := 0; i < height; i++ {
		world[i] = make([]bool, width)
	}

	return world
}
