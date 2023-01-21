package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/serjbibox/friend-service/pkg/models"
)

// Создание нового пользователя
func (h *Handler) createUser(c *gin.Context) {
	u := models.User{}
	err := c.BindJSON(&u)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := h.storage.User.Create(u)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, createResponse{
		ID: id,
	})
}

// Получение пользователя по ID
func (h *Handler) getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	u, err := h.storage.User.Get(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, u)

}

// Обновление данных пользователя по ID
func (h *Handler) updateUser(c *gin.Context) {
	u := models.User{}
	err := c.BindJSON(&u)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.storage.User.Update(u)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, messageResponse{
		Message: "ok",
	})
}

// Удаление данных пользователя по ID
func (h *Handler) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.storage.User.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, messageResponse{
		Message: "ok",
	})
}
