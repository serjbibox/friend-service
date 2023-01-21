package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/friend-service/pkg/models"
)

//Объект, реализующий интерфейс работы с таблицей posts PostgreSQL.
type AnketaPostgres struct {
	db  *pgxpool.Pool
	ctx context.Context
}

//Конструктор PostPostgres
func newAnketaPostgres(ctx context.Context, db *pgxpool.Pool) Anketa {
	return &AnketaPostgres{
		db:  db,
		ctx: ctx,
	}
}

// Получение публикаций по заданному количеству
func (s *AnketaPostgres) Get(id int) (models.Anketa, error) {
	return models.Anketa{}, nil
}

// Создание нового списка публикаций
func (s *AnketaPostgres) Create(u models.Anketa) (int, error) {
	return -1, nil
}

// Создание нового списка публикаций
func (s *AnketaPostgres) Update(u models.Anketa) error {
	return nil
}

// Получение публикаций по заданному количеству
func (s *AnketaPostgres) Delete(id int) error {
	return nil
}
