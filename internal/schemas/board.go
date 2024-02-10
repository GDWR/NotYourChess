package schemas

import "encoding/json"

type Board struct{}

func NewBoard() Board {
	return Board{}
}

func (b Board) ToString() string {
	return "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
}

func (b Board) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.ToString())
}
