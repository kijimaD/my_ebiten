package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kijimaD/my_ebiten/pkg/msg"
)

type Game struct{}

func (*Game) Update() error {
	return nil
}

func (*Game) Draw(screen *ebiten.Image) {
	msg.PrintMessage(screen, "hello")
}

func (*Game) Layout(int, int) (int, int) { return 50, 50 }

func main() {
	ebiten.SetWindowSize(128, 72)
	ebiten.SetInitFocused(false)
	ebiten.SetWindowTitle("Testing...")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
