package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/friend-service/pkg/models"
)

//Объект, реализующий интерфейс работы с таблицей posts PostgreSQL.
type SessionPostgres struct {
	db  *pgxpool.Pool
	ctx context.Context
}

//Конструктор PostPostgres
func newSessionPostgres(ctx context.Context, db *pgxpool.Pool) Session {
	return &SessionPostgres{
		db:  db,
		ctx: ctx,
	}
}

// Получение публикаций по заданному количеству
func (s *SessionPostgres) Get(id int) (models.Session, error) {
	return models.Session{}, nil
}

// Создание нового списка публикаций
func (s *SessionPostgres) Create(u models.Session) (int, error) {
	return -1, nil
}

// Создание нового списка публикаций
func (s *SessionPostgres) Update(u models.Session) error {
	return nil
}

// Получение публикаций по заданному количеству
func (s *SessionPostgres) Delete(id int) error {
	return nil
}
