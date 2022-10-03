package handler

import (
	"github.com/NEKETSKY/mnemosyne/pkg/log"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logger := log.LoggerFromGinContext(ctx)
	logger.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}
