package msg_test

import (
	"errors"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kijimaD/my_ebiten/pkg/msg"
)

const ImgDirPrefix = "../../test/actual/"

func TestExample_PrintMessage(t *testing.T) {
	const (
		Width  = 128
		Height = 72
	)

	tests := []struct {
		text string
	}{
		{text: "test"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("text_%v", i), func(t *testing.T) {
			e := os.Remove(ImgDirPrefix + tt.text + ".png")
			if e != nil {
				log.Fatal(e)
			}

			img := ebiten.NewImage(Width, Height)
			img.Fill(color.Black)

			file, err := os.Create(ImgDirPrefix + tt.text + ".png")
			defer file.Close()
			if err != nil {
				panic(err)
			}

			msg.PrintMessage(img, tt.text)
			if err := png.Encode(file, img); err != nil {
				panic(err)
			}
		})
	}
}

func TestMain(m *testing.M) {
	// hook
	RunTestGame(m)
}

// ================

var regularTermination = errors.New("regular termination")

type Game struct {
	m *testing.M
}

func (g *Game) Update() error {
	g.m.Run()
	// エラーを返さないと、実行終了しない
	return regularTermination
}

func (*Game) Draw(*ebiten.Image) {}

func (*Game) Layout(int, int) (int, int) { return 1, 1 }

func RunTestGame(m *testing.M) {
	ebiten.SetWindowSize(128, 72)
	ebiten.SetInitFocused(false)
	ebiten.SetWindowTitle("Testing...")

	g := &Game{
		m: m,
	}

	if err := ebiten.RunGame(g); err != nil && err != regularTermination {
		panic(err)
	}
	os.Exit(0)
}
