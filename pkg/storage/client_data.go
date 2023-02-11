package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/friend-service/pkg/models"
)

// Объект, реализующий интерфейс работы с таблицей posts PostgreSQL.
type ClientDataPostgres struct {
	db  *pgxpool.Pool
	ctx context.Context
}

// Конструктор PostPostgres
func newClientPostgres(ctx context.Context, db *pgxpool.Pool) ClientData {
	return &ClientDataPostgres{
		db:  db,
		ctx: ctx,
	}
}

// Получение списка анкет по заданному количеству
func (s *ClientDataPostgres) Get(id int) ([]models.ClientData, error) {
	return []models.ClientData{}, nil
}

// Создание нового списка публикаций
func (s *ClientDataPostgres) Create(u models.ClientData) (int, error) {
	return -1, nil
}

// Создание нового списка публикаций
func (s *ClientDataPostgres) Update(u models.ClientData) error {
	return nil
}

// Получение публикаций по заданному количеству
func (s *ClientDataPostgres) Delete(id int) error {
	return nil
}
