package handler

import (
	"github.com/ProjectUnion/project-backend.git/pkg/logging"
	"github.com/gin-gonic/gin"
)

type errorMsg struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger := logging.GetLogger()
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, errorMsg{message})
}
