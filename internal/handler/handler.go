package handler

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/internal/service"
	"github.com/NEKETSKY/mnemosyne/pkg/api"
)

type Handler struct {
	ctx      context.Context
	services service.Service
	api.MnemosyneServer
}

func NewHandler(ctx context.Context, services *service.Service) *Handler {
	return &Handler{
		ctx:      ctx,
		services: *services,
	}
}
