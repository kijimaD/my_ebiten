package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const ImgDir = "./test/actual/"

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
			img := ebiten.NewImage(Width, Height)
			img.Fill(color.Black)

			file, err := os.Create(ImgDir + tt.text + ".png")
			defer file.Close()
			if err != nil {
				panic(err)
			}

			PrintMessage(img, tt.text)
			if err := png.Encode(file, img); err != nil {
				panic(err)
			}
		})
	}
}

func TestMain(m *testing.M) {
	RunTestGame(m)
}

func PrintMessage(image *ebiten.Image, str string) {
	ebitenutil.DebugPrint(image, str)
}
