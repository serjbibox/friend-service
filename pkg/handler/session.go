package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/serjbibox/friend-service/pkg/models"
)

// Создание новой сессии
func (h *Handler) createSession(c *gin.Context) {
	s := models.Session{}
	err := c.BindJSON(&s)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := h.storage.Session.Create(s)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, createResponse{
		ID: id,
	})
}

// Получение сессии по ID
func (h *Handler) getSession(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	s, err := h.storage.Session.Get(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, s)
}

// Обновление данных сессии по ID
func (h *Handler) updateSession(c *gin.Context) {
	s := models.Session{}
	err := c.BindJSON(&s)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.storage.Session.Update(s)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, messageResponse{
		Message: "ok",
	})
}

// Удаление данных сессии по ID
func (h *Handler) deleteSession(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.storage.Session.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, messageResponse{
		Message: "ok",
	})
}
