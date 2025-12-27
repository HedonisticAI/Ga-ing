package helloworld

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SimpleGame struct {
	keys []ebiten.Key
}

func (g *SimpleGame) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}
	return nil
}

func (g *SimpleGame) Draw(screen *ebiten.Image) {

	for _, k := range g.keys {
		switch k {
		case ebiten.KeyT:
			ebitenutil.DebugPrint(screen, "Thai")
		case ebiten.KeyP:
			ebitenutil.DebugPrint(screen, "Partiot")
		default:
			ebitenutil.DebugPrint(screen, "Hello, World!")
		}
	}
}

func (g *SimpleGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
