package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/friend-service/pkg/models"
)

type User interface {
	Create(models.User) (int, error)
	Get(id int) (models.User, error)
	Update(models.User) error
	Delete(id int) error
}

type Session interface {
	Create(models.Session) (int, error)
	Get(id int) (models.Session, error)
	Update(models.Session) error
	Delete(id int) error
}

type Anketa interface {
	Create(models.Anketa) (int, error)
	Get(id int) (models.Anketa, error)
	Update(models.Anketa) error
	Delete(id int) error
}

type Storage struct {
	User
	Session
	Anketa
}

// Конструктор объекта хранилища для БД PostgreSQL.
func NewStoragePostgres(ctx context.Context, db *pgxpool.Pool) (*Storage, error) {
	if ctx == nil {
		return nil, errors.New("context is nil")
	}
	if db == nil {
		return nil, errors.New("db is nil")
	}
	return &Storage{
		User:    newUserPostgres(ctx, db),
		Session: newSessionPostgres(ctx, db),
		Anketa:  newAnketaPostgres(ctx, db),
	}, nil
}
