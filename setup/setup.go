package setup

import (
	"log"
	"github.com/BrianCarducci/go-game/game/characters"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func Setup() (characters.Player) {
	img, _, err := ebitenutil.NewImageFromFile("doggo.png")
	if err != nil {
		log.Fatal(err)
	}
	
	Player := characters.Player{
		Img: img,
		Xposition: 0,
		Yposition: 0,
	}

	return Player
}