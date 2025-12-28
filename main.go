package main

import (
	"log"
	helloworld "trying/HelloWorld"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Simple")
	if err := ebiten.RunGame(&helloworld.SimpleGame{}); err != nil {
		log.Fatal(err)
	}
}
