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

// Конструктор объекта Handler
func NewHandler(storage *storage.Storage) (*Handler, error) {
	if storage == nil {
		return nil, errors.New("storage is nil")
	}
	return &Handler{storage: storage}, nil
}

// Инициализация маршрутизатора запросов.
// Регистрация обработчиков запросов
func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	authMiddleware, err := h.newAuthMiddleWare()
	if err != nil {
		log.Fatalf("ошибка инициализации сервиса аутентификации: %s", err.Error())
	}

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.POST("/sign-up", h.identityUser)
	r.POST("/sign-in", authMiddleware.LoginHandler)
	user := r.Group("/user")
	user.GET("/refresh_token", authMiddleware.RefreshHandler)
	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.GET("/", h.getUser)
		user.PUT("/", h.updateUser)
		user.DELETE("/", h.deleteUser)
	}
	return r
}
