package handler

import (
	"errors"

	"github.com/gin-contrib/static"
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
	r.Use(static.Serve("/", static.LocalFile("./webapp", true)))
	r.GET("/news/:n", h.getNews)
	return r
}

// Получение публикаций по заданному количеству
func (h *Handler) getNews(c *gin.Context) {

}
