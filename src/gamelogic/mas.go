package gamelogic

import (
	"errors"
	"math/rand"
	"time"
)

type Mas struct {
	Deck          []Card
	Discarddeck   []Card
	Playdeck      []Card
	Players       map[string][]Card
	Playerorder   []string
	State         int
	PlayerTracker int
}

func NewMas() *Mas {
	mas := Mas{make([]Card, 52), make([]Card, 0, 52), make([]Card, 0, 52), make(map[string][]Card, 52/3), make([]string, 0, 52/3), 0, 0}
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
		return errors.New("can't add players in this state of the game")
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
	self.PlayerTracker = rand.Intn(len(self.Playerorder))
	self.State = 1
}

func (self *Mas) Turn(card Card) (*string, error) {
	return nil, nil //Not implemented
}

func (self *Mas) cardFromPlayerToDeck(player string, deck *[]Card, card Card) error {
	playerdeck := self.Players[player]
	if len(playerdeck) == 0 {
		return errors.New("player doesn't exist/has empty hand")
	}
	index := linear_search(card, playerdeck)
	if index < 0 {
		return errors.New("card not in deck")
	}
	delete(index, &playerdeck)
	self.Playdeck = append(*deck, card)
	return nil
}

func (self *Mas) gatherRound(card Card) (*string, error) {
	[]func(){
		func() (*string, error) {
			e := self.cardFromPlayerToDeck(self.Playerorder[self.PlayerTracker], &self.Playdeck, card)
			if e != nil {
				return nil, e
			}
			self.PlayerTracker++
			return self.Playerorder[self.PlayerTracker], nil
		},
		func() (*string, error) {
			e := self.cardFromPlayerToDeck(self.Playerorder[self.PlayerTracker], &self.Playdeck, card)
			if e != nil {
				return nil, e
			}
			if self.Playdeck[-1].value > self.Playdeck[-2].value { // Last player took the hand
				
			}
		}
	}
}

func linear_search(card Card, lst []Card) int {
	for i := 0; i < len(lst); i++ {
		if card.value == lst[i].value && card.valor == card.valor {
			return i
		}
	}
	return -1
}

func delete(index int, lst *[]Card) {
	*lst = append((*lst)[:index], (*lst)[index:]...)
}
