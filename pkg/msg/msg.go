package msg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func PrintMessage(image *ebiten.Image, str string) {
	ebitenutil.DebugPrint(image, str)
}
