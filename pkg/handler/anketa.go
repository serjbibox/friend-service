package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/serjbibox/friend-service/pkg/models"
)

// Создание нового пользователя
func (h *Handler) createAnketa(c *gin.Context) {
	a := models.Anketa{}
	err := c.BindJSON(&a)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := h.storage.Anketa.Create(a)
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
func (h *Handler) getAnketa(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	a, err := h.storage.Anketa.Get(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, a)

}

// Обновление данных пользователя по ID
func (h *Handler) updateAnketa(c *gin.Context) {
	a := models.Anketa{}
	err := c.BindJSON(&a)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.storage.Anketa.Update(a)
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
func (h *Handler) deleteAnketa(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.storage.Anketa.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, messageResponse{
		Message: "ok",
	})
}
