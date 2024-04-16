package database

import (
	"errors"
	. "pointsCounterCRUD/database/model"
)

var (
	exampleError = errors.New("example")
)

type Repository interface {
	Migrate() error
	Login()

	CreateScore(User, Stat) (*Stat, error)
	GetScores(User) ([]Stat, error)
	ErraceScore(Stat) (*Stat, error)

	CreateUser(User) (*User, error)
}
