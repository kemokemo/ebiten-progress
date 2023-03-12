package main

import (
	"image/color"
	_ "image/png" // to load png images

	"github.com/hajimehoshi/ebiten/v2"
	prog "local.packages/progress"
)

const (
	screenWidth  = 480
	screenHeight = 320
)

func NewGame() *Game {
	img := ebiten.NewImage(20, 20)
	img.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	return &Game{progress: prog.NewProgress(img, 40.0, 5.0, 400.0)}
}

type Game struct {
	progress *prog.Progress
	value    int
	counter  int
}

func (g *Game) Update() error {
	g.counter++
	if g.counter > 10 {
		g.counter = 0
		g.value++
	}

	g.progress.SetPercent(g.value)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.progress.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
