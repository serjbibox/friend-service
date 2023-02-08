package handler

import (
	"errors"
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/serjbibox/friend-service/pkg/storage"
)

// Обработчик HTTP запросов сервера GoNews
type Handler struct {
	storage *storage.Storage
}

//Конструктор объекта Handler
func NewHandler(storage *storage.Storage) (*Handler, error) {
	if storage == nil {
		return nil, errors.New("storage is nil")
	}
	return &Handler{storage: storage}, nil
}

//Инициализация маршрутизатора запросов.
//Регистрация обработчиков запросов
func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	authMiddleware, err := newAuthMiddleWare()
	if err != nil {
		log.Fatalf("ошибка инициализации сервиса аутентификации: %s", err.Error())
	}
	r.POST("/login", authMiddleware.LoginHandler)
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
	}

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
