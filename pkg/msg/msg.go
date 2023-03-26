package msg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// var regularTermination = errors.New("regular termination")

// type Game struct {
// 	m *testing.M
// }

// func (g *Game) Update(*ebiten.Image) error {
// 	g.m.Run()
// 	return regularTermination
// }

// func (*Game) Draw(*ebiten.Image) {}

// func (g *Game) Layout(int, int) (int, int) { return 1, 1 }

// func RunTestGame(m *testing.M) {
// 	ebiten.SetWindowSize(128, 72)
// 	ebiten.SetInitFocused(false)
// 	ebiten.SetWindowTitle("Testing...")

// 	g := &Game{
// 		m: m,
// 	}

// 	if err := ebiten.RunGame(g); err != nil && err != regularTermination {
// 		panic(err)
// 	}
// 	os.Exit(0)
// }

func PrintMessage(image *ebiten.Image, str string) {
	ebitenutil.DebugPrint(image, str)
}
