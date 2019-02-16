package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type game struct {
	width  int
	height int
	cells  [][]bool
}

func newCells(width, height int) [][]bool {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}

	return cells
}

func New(width, height int) *game {
	g := &game{
		width:  width,
		height: height,
		cells:  newCells(width, height),
	}
	return g
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
	for _, rowCells := range g.cells {
		for _, cell := range rowCells {
			if cell {
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println("")
	}
}

func (g *game) update() {
	newCell := newCells(g.width, g.height)
	// TODO: implement here!

	g.cells = newCell
}

func main() {
	g := New(50, 20)
	for {
		g.showCells()
		g.update()
		time.Sleep(1 * time.Second)
	}
}
