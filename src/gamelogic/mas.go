package gamelogic

import (
	"errors"
	"math/rand"
	"time"
)

type Mas struct {
	Deck        []Card
	Discarddeck []Card
	Playdeck    []Card
	Players     map[string][]Card
	Playerorder []string
	State       int
}

func NewMas() *Mas {
	mas := Mas{make([]Card, 52, 52), make([]Card, 0, 52), make([]Card, 0, 52), make(map[string][]Card, 52/3), make([]string, 0, 52/3), 0}
	for i := 0; i < 52; i++ {
		mas.Deck[i] = Card{i/13 + 1, i%13 + 1} // Populate deck with all types of cards
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(mas.Deck), func(i, j int) { mas.Deck[i], mas.Deck[j] = mas.Deck[j], mas.Deck[i] })
	return &mas
}

func NewMasWithPlayers(names []string) (*Mas, error) {
	mas := NewMas()
	for _, name := range names {
		panicq := mas.AddPlayer(name)
		if panicq != nil {
			return nil, panicq
		}
	}
	return mas, nil
}

func (self *Mas) AddPlayer(name string) error {
	if self.State != 0 {
		return errors.New("Can't add players in this state of the game")
	}
	self.Players[name] = make([]Card, 3, 52)
	for i := 0; i < 3; i++ {
		self.Players[name][i] = self.Deck[i]
	}
	self.Deck = self.Deck[3:]
	self.Playerorder = append(self.Playerorder, name)
	return nil
}

func (self *Mas) Start() {
	self.State = 1
}
