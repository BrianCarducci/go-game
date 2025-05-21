package setup

import (
	"bytes"
	"embed"
	"log"

	"github.com/BrianCarducci/go-game/game/characters"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
)


func Setup(fontFiles embed.FS) (characters.Player, *text.GoTextFace) {
	defaultFontFile, err := fontFiles.ReadFile("assets/fonts/Montserrat-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}

	fontFaceSource, err := text.NewGoTextFaceSource(bytes.NewReader(defaultFontFile)) 

	fontFace := &text.GoTextFace{
		Source: fontFaceSource,
		Direction: text.DirectionLeftToRight,
		Size: 24,
		Language: language.AmericanEnglish,
	}

	if err != nil {
		log.Fatal(err)
	}

	img, _, err := ebitenutil.NewImageFromFile("assets/sprites/doggo.png")
	if err != nil {
		log.Fatal(err)
	}
	
	Player := characters.Player{
		Img: img,
		Xposition: 0,
		Yposition: 0,
	}

	return Player, fontFace
}