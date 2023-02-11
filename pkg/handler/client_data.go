package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/serjbibox/friend-service/pkg/models"
)

// Создание нового пользователя
func (h *Handler) createClientData(c *gin.Context) {
	a := models.ClientData{}
	err := c.BindJSON(&a)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := h.storage.ClientData.Create(a)
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
func (h *Handler) getClientData(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	a, err := h.storage.ClientData.Get(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, a)

}

// Обновление данных пользователя по ID
func (h *Handler) updateClientData(c *gin.Context) {
	a := models.ClientData{}
	err := c.BindJSON(&a)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.storage.ClientData.Update(a)
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
func (h *Handler) deleteClientData(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.storage.ClientData.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, messageResponse{
		Message: "ok",
	})
}
