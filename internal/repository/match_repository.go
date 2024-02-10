package repository

import "github.com/gdwr/chaoss/internal/schemas"

type MatchRepository interface {
	RandomMatch() (*schemas.Match, error)
	NewMatch() *schemas.Match
	GetMatch(id schemas.Guid) (*schemas.Match, error)
}
