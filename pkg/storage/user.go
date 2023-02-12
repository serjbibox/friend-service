package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/friend-service/pkg/models"
)

// Объект, реализующий интерфейс работы с таблицей posts PostgreSQL.
type UserPostgres struct {
	db  *pgxpool.Pool
	ctx context.Context
}

// Конструктор PostPostgres
func newUserPostgres(ctx context.Context, db *pgxpool.Pool) User {
	return &UserPostgres{
		db:  db,
		ctx: ctx,
	}
}

// Получение таблицы User по id
func (s *UserPostgres) GetById(id int) (models.User, error) {

	return models.User{}, nil
}

// Получение таблицы User по номеру телефона
func (s *UserPostgres) GetByPhone(phone string) (models.User, error) {

	return models.User{}, nil
}

// Создание нового списка публикаций
func (s *UserPostgres) Create(u models.User) (int, error) {
	fmt.Println("Сделаль!")
	return -1, nil
}

// Создание нового списка публикаций
func (s *UserPostgres) Update(u models.User) error {
	return nil
}

// Получение публикаций по заданному количеству
func (s *UserPostgres) Delete(id int) error {
	return nil
}
