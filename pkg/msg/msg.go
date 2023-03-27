package msg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func PrintMessage(image *ebiten.Image, str string) {
	ebitenutil.DebugPrint(image, str)
}

type Game struct{}

func (g *Game) Update() error { return nil }

func (*Game) Draw(i *ebiten.Image) {
	ebitenutil.DebugPrint(i, "hello")
}

func (*Game) Layout(int, int) (int, int) { return 40, 40 }

func RunGame() {
	ebiten.SetWindowSize(128, 72)
	ebiten.SetInitFocused(false)
	ebiten.SetWindowTitle("Testing...")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
