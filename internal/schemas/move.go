package schemas

import (
	"encoding/json"
	"fmt"
)

type Move struct {
	Piece string
	From  string
	To    string
}

func NewMoveFromString(s string) Move {
	return Move{
		Piece: string(s[0]),
		From:  string(s[1:3]),
		To:    string(s[3:5]),
	}
}

func (m Move) ToString() string {
	return fmt.Sprintf("%s%s%s", m.Piece, m.From, m.To)
}

func (m Move) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.ToString())
}
