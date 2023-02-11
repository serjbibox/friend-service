package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/friend-service/pkg/models"
)

// Объект, реализующий интерфейс работы с таблицей posts PostgreSQL.
type PsyDataPostgres struct {
	db  *pgxpool.Pool
	ctx context.Context
}

// Конструктор PostPostgres
func newPsyPostgres(ctx context.Context, db *pgxpool.Pool) PsyData {
	return &PsyDataPostgres{
		db:  db,
		ctx: ctx,
	}
}

// Получение списка анкет по заданному количеству
func (s *PsyDataPostgres) Get(id int) ([]models.PsyData, error) {
	return []models.PsyData{}, nil
}

// Создание нового списка публикаций
func (s *PsyDataPostgres) Create(u models.PsyData) (int, error) {
	return -1, nil
}

// Создание нового списка публикаций
func (s *PsyDataPostgres) Update(u models.PsyData) error {
	return nil
}

// Получение публикаций по заданному количеству
func (s *PsyDataPostgres) Delete(id int) error {
	return nil
}
