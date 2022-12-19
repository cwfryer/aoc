package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/luisparravicini/gtris/gtris"
)

func getInstruction(fn string) []ebiten.Key {
    file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

    keyMap := make(map[rune]ebiten.Key)
    keyMap['<'] = ebiten.KeyLeft
    keyMap['>'] = ebiten.KeyRight
	scanner := bufio.NewScanner(file)
    var instructions []ebiten.Key
	for scanner.Scan() {
        for _,c := range scanner.Text() {
            km := keyMap[c]
            instructions = append(instructions,km)
        }
	}
    return instructions
}

func main() {
    i := getInstruction("../day17.input")
	rand.Seed(time.Now().UnixNano())

	game := gtris.NewGame(i)

	ebiten.SetWindowSize(gtris.ScreenWidth, gtris.ScreenHeight)
	ebiten.SetWindowTitle("gtris")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
