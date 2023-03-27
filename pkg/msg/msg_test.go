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

// 画像をうまく出力するためには、ウィンドウを出した状態でテストを実行する必要がある
// beforeフックでゲームを実行しつつ、画像をファイルに書き出すテストを実行する
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
	RunTestGame(m)
}

// ================

var regularTermination = errors.New("regular termination")

type TestGame struct {
	msg.Game
	m *testing.M
}

func (g *TestGame) Update() error {
	// m.Run()はテストの前に実行される
	g.m.Run()
	// エラーを返さないと、実行終了しない
	return regularTermination
}

func RunTestGame(m *testing.M) {
	ebiten.SetWindowSize(128, 72)
	ebiten.SetInitFocused(false)
	ebiten.SetWindowTitle("Testing...")

	g := &TestGame{
		m: m,
	}

	if err := ebiten.RunGame(g); err != nil && err != regularTermination {
		panic(err)
	}
	os.Exit(0)
}
