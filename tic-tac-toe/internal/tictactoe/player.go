package tictactoe

type Player struct {
	Name string
	Symbol rune
}

func NewPlayer(Name string, Symbol rune) *Player {
	return &Player{Name: Name, Symbol: Symbol}
}
