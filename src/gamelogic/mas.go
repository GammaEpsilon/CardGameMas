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
	Players       map[string][2][]Card //TODO Fix
	Playerorder   []string
	State         int
	PlayerTracker int
}

func NewMas() *Mas {
	mas := Mas{make([]Card, 52), make([]Card, 0, 52), make([]Card, 0, 52), make(map[string][2][]Card, 52/3), make([]string, 0, 52/3), 0, 0}
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
	self.Players[name] = [2][]Card{make([]Card, 3, 52), make([]Card, 0, 52)}
	for i := 0; i < 3; i++ {
		self.Players[name][0][i] = self.Deck[i]
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

func (self *Mas) cardFromDeckToDeck(from *[]Card, to *[]Card, card Card) error {
	index := linear_search(card, from)
	if index < 0 {
		return errors.New("card not in deck")
	}
	delete(index, from)
	*to = append(*to, card)
	return nil
}

func (self *Mas) cardFromPlayerToDeck(player string, deck *[]Card, card Card) error {
	playerdeck := self.Players[player][0]
	if len(playerdeck) == 0 {
		return errors.New("player doesn't exist/has empty hand")
	}
	return self.cardFromDeckToDeck(&playerdeck, deck, card)
}

func (self *Mas) gatherRound(card Card) (*string, error) {
	return []func() (*string, error){
		func() (*string, error) {
			e := self.cardFromPlayerToDeck(self.Playerorder[self.PlayerTracker], &self.Playdeck, card)
			if e != nil {
				return nil, e
			}
			self.Players[self.Playerorder[self.PlayerTracker]][0] = append(self.Players[self.Playerorder[self.PlayerTracker]], self.Deck[1])
			self.Deck = self.Deck[1:]
			self.PlayerTracker++
			return &self.Playerorder[self.PlayerTracker], nil
		},
		func() (*string, error) {
			e := self.cardFromPlayerToDeck(self.Playerorder[self.PlayerTracker], &self.Playdeck, card)
			if e != nil {
				return nil, e
			}
			self.Players[self.Playerorder[self.PlayerTracker]] = append(self.Players[self.Playerorder[self.PlayerTracker]], self.Deck[1])
			self.Deck = self.Deck[1:]
			if self.Playdeck[len(self.Playdeck)-1].value <= self.Playdeck[len(self.Playdeck)-2].value { // Last player did not take the hand
				self.PlayerTracker-- //backtrack tracker
			}
			if self.Playdeck[len(self.Playdeck)-1].value != self.Playdeck[len(self.Playdeck)-2].value { // A player took the hand
				self.Players[self.Playerorder[self.PlayerTracker]] = append(self.Players[self.Playerorder[self.PlayerTracker]], self.Playdeck...) // Add all cards to player
				self.Playdeck = self.Playdeck[:0]                                                                                                 // Empty playdeck
			} // If no player took the hand, keep playing for it
			return &self.Playerorder[self.PlayerTracker], nil
		},
	}[len(self.Playdeck)%2]() //State machine implementation
}

func linear_search(card Card, lst *[]Card) int {
	for i := 0; i < len(*lst); i++ {
		if card.value == (*lst)[i].value && card.valor == card.valor {
			return i
		}
	}
	return -1
}

func delete(index int, lst *[]Card) {
	*lst = append((*lst)[:index], (*lst)[index:]...)
}
