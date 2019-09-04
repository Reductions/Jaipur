package server

import (
	"math/rand"
	"time"

	. "github.com/Reductions/jaipur/common"
)

var seeded bool = false

var numberOfCards = map[Good]int{
	Camel:   8,
	Leather: 10,
	Cloth:   8,
	Spice:   8,
	Silver:  6,
	Gold:    6,
	Diamond: 6,
}

var initialGoodsTokenSupply = map[Good][]Token{
	Camel:   []Token{5},
	Leather: []Token{4, 3, 2, 1, 1, 1, 1, 1, 1},
	Cloth:   []Token{5, 3, 3, 2, 2, 1, 1},
	Spice:   []Token{5, 3, 3, 2, 2, 1, 1},
	Silver:  []Token{5, 5, 5, 5, 5},
	Gold:    []Token{6, 6, 5, 5, 5},
	Diamond: []Token{7, 7, 5, 5, 5},
}

var (
	Bonus3Tokens = []Token{3, 3, 3, 2, 2, 1, 1}
	Bonus4Tokens = []Token{6, 6, 5, 5, 4, 4}
	Bonus5Tokens = []Token{10, 10, 9, 8, 8}
)

func seedIfNotSeeded() {
	if seeded {
		return
	}
	rand.Seed(time.Now().UTC().UnixNano())
	seeded = true
}

func newDeck() Deck {
	result := make([]Good, DeckSize)
	iter := 0
	for ; iter < 3; iter++ {
		result[iter] = Camel
	}
	for good, count := range numberOfCards {
		for i := 0; i < count; i++ {
			result[iter] = good
			iter++
		}
	}
	return result
}

func NewShuffledDeck() Deck {
	deck := newDeck()
	seedIfNotSeeded()
	rand.Shuffle(len(deck), func(i, j int) {
		if i < 3 || j < 3 {
			return
		}
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

func goodsTokenSupply() map[Good][]Token {
}

func NewTokenSupply() TokenSupply {
	supply := new(TokenSupply)
	supply.GoodTokens = initialGoodsTokenSupply
	supply.GoodTokens
	return supply
}
