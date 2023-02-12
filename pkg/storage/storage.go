package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/friend-service/pkg/models"
)

type User interface {
	Create(models.User) (int, error)
	GetById(id int) (models.User, error)
	GetByPhone(phone string) (models.User, error)
	Update(models.User) error
	Delete(id int) error
}

type Session interface {
	Create(models.Session) (int, error)
	Get(id int) (models.Session, error)
	Update(models.Session) error
	Delete(id int) error
}

type PsyData interface {
	Create(models.PsyData) (int, error)
	Get(id int) ([]models.PsyData, error)
	Update(models.PsyData) error
	Delete(id int) error
}

type ClientData interface {
	Create(models.ClientData) (int, error)
	Get(id int) ([]models.ClientData, error)
	Update(models.ClientData) error
	Delete(id int) error
}

type Storage struct {
	User
	Session
	PsyData
	ClientData
	Cache map[string]int
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
		User:       newUserPostgres(ctx, db),
		Session:    newSessionPostgres(ctx, db),
		PsyData:    newPsyPostgres(ctx, db),
		ClientData: newClientPostgres(ctx, db),
		Cache:      make(map[string]int),
	}, nil
}
