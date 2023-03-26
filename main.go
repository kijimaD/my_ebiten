package main

import (
	"errors"
	"os"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

var regularTermination = errors.New("regular termination")

type game struct {
	m    *testing.M
	code int
}

func (g *game) Update() error {
	g.code = g.m.Run()
	return regularTermination
}

func (*game) Draw(*ebiten.Image) {}

func (g *game) Layout(int, int) (int, int) { return 1, 1 }

func RunTestGame(m *testing.M) {
	ebiten.SetWindowSize(128, 72)
	ebiten.SetInitFocused(false)
	ebiten.SetWindowTitle("Testing...")

	g := &game{
		m: m,
	}
	if err := ebiten.RunGame(g); err != nil && err != regularTermination {
		panic(err)
	}
	os.Exit(g.code)
}
