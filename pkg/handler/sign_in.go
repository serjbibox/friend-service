package handler

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/serjbibox/friend-service/pkg/models"
)

// Аутентификация пользователя
func (h *Handler) signIn(c *gin.Context) (interface{}, error) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	phone := loginVals.Phone
	password := loginVals.Password
	role := loginVals.Role
	if (phone == "996550269716") && (password == "123456") {
		u := &models.User{
			Phone: phone,
			Role:  role,
		}
		return u, nil
	}

	return nil, jwt.ErrFailedAuthentication
}
