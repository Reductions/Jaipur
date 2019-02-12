package common

type Good byte

const (
	Camel Good = iota
	Leather
	Cloth
	Spice
	Silver
	Gold
	Diamond
)

const DeckSize byte = 55

type Deck []Good

type Token byte

type TokenSupply struct {
	GoodTokens   map[Good][]Token
	Bonus3Tokens []Token
	Bonus4Tokens []Token
	Bonus5Tokens []Token
}
