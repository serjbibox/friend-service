package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/friend-service/pkg/models"
	"golang.org/x/crypto/bcrypt"
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

// Создание нового пользователя
func (s *UserPostgres) Create(u models.User) (int, error) {
	id := 0
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	query := `
	INSERT INTO users(phone, role, password)
	VALUES ($1, $2, $3) RETURNING ID;`
	err = s.db.QueryRow(context.Background(), query, u.Phone, u.Role, hashedPassword).Scan(&id)
	log.Println(id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// Обновление пользователя
func (s *UserPostgres) Update(u models.User) error {
	return nil
}

// Удаление пользователя
func (s *UserPostgres) Delete(id int) error {
	return nil
}
