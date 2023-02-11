package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/serjbibox/friend-service/pkg/models"
)

const (
	smsCodeMin   = 100000
	smsCodeMax   = 999999
	smsSent      = "sms sent, code: "
	wrongCode    = "wrong sms code"
	wrongPhone   = "wrong phone"
	client       = "client"
	psychologist = "psy"
	moderator    = "mod"
)

func (h *Handler) sendSms(p string) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	code := smsCodeMin + rand.Intn(smsCodeMax-smsCodeMin+1)
	h.storage.Cache[p] = code
	return code
}

// Создание нового пользователя
func (h *Handler) identityUser(c *gin.Context) {
	u := models.UserIdentity{}
	err := c.BindJSON(&u)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	if u.SmsCode == 0 {
		code := h.sendSms(u.Phone)
		c.JSON(http.StatusOK, identityResponse{
			Phone:   u.Phone,
			Message: smsSent + fmt.Sprint(code),
		})
		return
	}
	if _, ok := h.storage.Cache[u.Phone]; !ok {
		c.JSON(http.StatusOK, identityResponse{
			Phone:   u.Phone,
			Message: wrongPhone,
		})
		return
	}
	if u.SmsCode != h.storage.Cache[u.Phone] {
		c.JSON(http.StatusOK, identityResponse{
			Phone:   u.Phone,
			Message: wrongCode,
		})
		return
	}
	delete(h.storage.Cache, u.Phone)
	switch u.Role {
	case psychologist:
	case moderator:
	case client:
		fallthrough
	default:
		user := models.User{}
		user.Phone = u.Phone
		user.Role = client
		user.Password = u.Password
		id, err := h.storage.User.Create(user)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, createResponse{
			ID: id,
		})
	}

}
