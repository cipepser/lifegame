package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type game struct {
	width      int
	height     int
	cells      [][]bool
	generation int
}

func newCells(width, height int) [][]bool {
	cells := make([][]bool, height+1)
	for i := range cells {
		cells[i] = make([]bool, width+1)
	}

	return cells
}

func New(width, height int) *game {
	g := &game{
		width:      width,
		height:     height,
		cells:      newCells(width, height),
		generation: 0,
	}
	return g
}

func (g *game) initialize() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < g.height-1; i++ {
		for j := 1; j < g.width-1; j++ {
			if rand.Int()%2 == 0 {
				g.cells[i][j] = true
			}
		}
		fmt.Println("")
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func (g *game) showCells() {
	clear()
	for i := 1; i < g.height-1; i++ {
		for j := 1; j < g.width-1; j++ {
			if g.cells[i][j] {
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println("")
	}
	fmt.Println("generation: ", g.generation)
}

func (g *game) update() {
	newCell := newCells(g.width, g.height)
	for i := 1; i < g.height-1; i++ {
		for j := 1; j < g.width-1; j++ {
			cnt := g.countLivingNeighbours(i, j)
			if g.cells[i][j] {
				if cnt == 2 || cnt == 3 {
					newCell[i][j] = true
				}
			} else {
				if cnt == 3 {
					newCell[i][j] = true
				}
			}
		}
	}

	g.cells = newCell
	g.generation++
}

func (g *game) countLivingNeighbours(i, j int) int {
	cnt := 0
	for k := -1; k <= 1; k++ {
		for l := -1; l <= 1; l++ {
			if k == 0 && l == 0 {
				continue
			}
			if g.cells[i+k][j+l] {
				cnt++
			}
		}
	}

	return cnt
}

func main() {
	g := New(50, 20)
	g.initialize()
	for {
		g.showCells()
		g.update()
		time.Sleep(200 * time.Millisecond)
	}
}
