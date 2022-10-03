package handler

import (
	"context"
	_ "github.com/NEKETSKY/mnemosyne/docs"
	"github.com/NEKETSKY/mnemosyne/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	ctx      context.Context
	services service.Service
}

func NewHandler(ctx context.Context, services *service.Service) *Handler {
	return &Handler{
		ctx:      ctx,
		services: *services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	list := router.Group("/mnemosyne", h.logger)
	{
		list.POST("/", h.mnemosyneRequest)
	}

	return router
}
