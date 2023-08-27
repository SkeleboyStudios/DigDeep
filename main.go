package main

import (
	"github.com/EngoEngine/engo"

	"github.com/SkeleboyStudios/DigDeep/scenes"
)

func main() {
	engo.RegisterScene(&scenes.MainMenuScene{})
	engo.Run(engo.RunOptions{
		Title:         "Dig Deeper",
		Width:         640,
		Height:        360,
		ScaleOnResize: true,
		//}, &scenes.MadeScene{})
	}, &scenes.MainMenuScene{})
}
