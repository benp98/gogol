package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"time"

	"github.com/benp98/gogol"
	"golang.org/x/exp/rand"
)

var generations int
var width int
var height int
var neighbourRadius int

var palette = []color.Color{
	color.Gray{255},
	color.Gray{0},
}

func main() {
	flag.IntVar(&generations, "g", 100, "generation count")
	flag.IntVar(&width, "w", 128, "width of the world")
	flag.IntVar(&height, "h", 128, "height of the world")
	flag.IntVar(&neighbourRadius, "n", 1, "neighbour radius")
	flag.Parse()

	outfile := flag.Arg(0)

	// Check if the argument is supplied
	if len(outfile) == 0 {
		fmt.Println("Usage: gogol [optional flags] <outfile.gif>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if neighbourRadius < 1 {
		fmt.Println("Neighbour radius must be 1 or greater")
		flag.PrintDefaults()
		os.Exit(1)
	}

	state := gogol.NewState(width, height, neighbourRadius)

	randomizeWorld(state)

	//state.SetCell(10, 10, true)
	//state.SetCell(11, 10, true)
	//state.SetCell(12, 10, true)
	//state.SetCell(12, 9, true)
	//state.SetCell(11, 8, true)

	gifData := new(gif.GIF)
	gifData.LoopCount = 0

	for i := 0; i < generations; i++ {
		gifData.Image = append(gifData.Image, renderWorld(state))
		gifData.Delay = append(gifData.Delay, 1)
		state.NextGeneration()
	}

	file, err := os.Create(outfile)
	defer file.Sync()
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gif.EncodeAll(file, gifData)
}

func randomizeWorld(state *gogol.State) {
	rand.Seed(uint64(time.Now().Unix()))

	modulo := (neighbourRadius * neighbourRadius) + 1

	width, height := state.GetDimensions()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			state.SetCell(x, y, rand.Int()%modulo == 0)
		}
	}
}

func renderWorld(state *gogol.State) *image.Paletted {
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)

	width, height := state.GetDimensions()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := state.GetCell(x, y)

			if cell {
				img.SetColorIndex(x, y, 1)
			} else {
				img.SetColorIndex(x, y, 0)
			}
		}
	}

	return img
}
