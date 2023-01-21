package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/serjbibox/friend-service/pkg/storage"
)

// Обработчик HTTP запросов сервера GoNews
type Handler struct {
	storage *storage.Storage
}

//Конструктор объекта Handler
func New(storage *storage.Storage) (*Handler, error) {
	if storage == nil {
		return nil, errors.New("storage is nil")
	}
	return &Handler{storage: storage}, nil
}

//Инициализация маршрутизатора запросов.
//Регистрация обработчиков запросов
func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/user", h.getUser)
	r.GET("/session", h.getSession)
	r.GET("/anketa", h.getAnketa)
	r.POST("/user", h.createUser)
	r.POST("/session", h.createSession)
	r.POST("/anketa", h.createAnketa)
	r.PUT("/user", h.updateUser)
	r.PUT("/session", h.updateSession)
	r.PUT("/anketa", h.updateAnketa)
	r.DELETE("/user", h.deleteUser)
	r.DELETE("/session", h.deleteSession)
	r.DELETE("/anketa", h.deleteAnketa)
	return r
}
