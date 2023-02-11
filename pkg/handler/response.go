package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type messageResponse struct {
	Message string `json:"message"`
}

type identityResponse struct {
	Phone   string `json:"phone" example:"+79161234567"`
	Message string `json:"message" example:"idented"`
}

type createResponse struct {
	ID int `json:"id"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
