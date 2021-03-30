package main

import (
	"fmt"
	"strings"

	"github.com/GammaEpsilon/CardGameMas/src/gamelogic"
)

func main() {
	mas, _ := gamelogic.NewMasWithPlayers([]string{"Erik", "Sara", "Mats", "Mikael"})
	for key, val := range mas.Players {
		strlst := make([]string, 0, len(val))
		for _, card := range val {
			strlst = append(strlst, card.CardToString())
		}
		fmt.Printf("%s: %s\n", key, strings.Join(strlst, ", "))
	}
}
