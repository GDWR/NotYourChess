package repository

import (
	"errors"
	"math/rand"

	"github.com/gdwr/chaoss/internal/schemas"
)

type InMemoryMatchRepository struct {
	matches []schemas.Match
}

func NewInMemoryMatchRepository() InMemoryMatchRepository {
	return InMemoryMatchRepository{}
}

func (r InMemoryMatchRepository) RandomMatch() (*schemas.Match, error) {
	if len(r.matches) == 0 {
		return nil, errors.New("no matches")
	}

	index := rand.Intn(len(r.matches))
	return &r.matches[index], nil
}

func (r InMemoryMatchRepository) NewMatch() *schemas.Match {
	newMatch := schemas.Match{
		Id:    schemas.NewGuid(),
		Board: schemas.NewBoard(),
		Moves: []schemas.Move{},
	}

	r.matches = append(r.matches, newMatch)
	return &newMatch
}
