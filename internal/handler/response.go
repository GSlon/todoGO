package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func errorResponse(c *gin.Context, statusCode int, message string) {
    logrus.Error(message)
    c.AbortWithStatusJSON(statusCode, ResponseError{message})
}
