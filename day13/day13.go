package main

import (
	"fmt"
	"github.com/dhconnelly/advent-of-code-2019/breakout"
	"github.com/dhconnelly/advent-of-code-2019/intcode"
	"log"
	"os"
)

func countTiles(
	state breakout.GameState,
	which breakout.TileId,
) int {
	n := 0
	for _, tile := range state.Tiles {
		if tile == which {
			n++
		}
	}
	return n
}

func main() {
	data, err := intcode.ReadProgram(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	state, err := breakout.Play(data, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(countTiles(state, breakout.BLOCK))
}