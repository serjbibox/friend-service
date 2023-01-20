package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/friend-service/pkg/models"
)

//Объект, реализующий интерфейс работы с таблицей posts PostgreSQL.
type UserPostgres struct {
	db  *pgxpool.Pool
	ctx context.Context
}

//Конструктор PostPostgres
func newUserPostgres(ctx context.Context, db *pgxpool.Pool) User {
	return &UserPostgres{
		db:  db,
		ctx: ctx,
	}
}

// Получение публикаций по заданному количеству
func (s *UserPostgres) Get(id int) (models.User, error) {
	return models.User{}, nil
}

// Создание нового списка публикаций
func (s *UserPostgres) Create(u models.User) error {
	return nil
}

// Создание нового списка публикаций
func (s *UserPostgres) Update(u models.User) error {
	return nil
}

// Получение публикаций по заданному количеству
func (s *UserPostgres) Delete(id int) error {
	return nil
}
