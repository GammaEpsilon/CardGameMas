package gamelogic

import (
	"errors"
	"strconv"
)

var valor_map = map[string]int{
	"clover":   1,
	"diamonds": 2,
	"spade":    3,
	"hearts":   4,
}

type Card struct {
	valor int //1 to 4 represents clover diamonds spade hearts respectively
	value int //1-13
}

func NewCard(valor string, value int) (*Card, error) {
	if value < 1 || value > 13 {
		return nil, errors.New("value not in range 1-13")
	}
	intrep := valor_map[valor]
	if intrep == 0 {
		var valstr string
		for val := range valor_map {
			valstr += val + " "
			return nil, errors.New("invalid valor. valid valors are:" + valstr)
		}
	}
	return &Card{intrep, value}, nil
}

func (self *Card) CardToString() string {
	var str string
	if self.value > 10 {
		str = []string{"Knight", "Queen", "King", "Ace"}[self.value-11]
	} else {
		str = strconv.Itoa(self.value)
	}
	for key, val := range valor_map {
		if val == self.valor {
			return str + " of " + key
		}
	}
	return str + " " + strconv.Itoa(self.valor)
}
