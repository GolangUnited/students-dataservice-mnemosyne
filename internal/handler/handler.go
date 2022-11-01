package handler

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/internal/service"
	"github.com/NEKETSKY/mnemosyne/pkg/api"
)

// Handler struct with grpc api server
type Handler struct {
	ctx      context.Context
	services service.Service
	api.MnemosyneServer
}

// NewHandler creates a new handler
func NewHandler(ctx context.Context, services *service.Service) *Handler {
	return &Handler{
		ctx:      ctx,
		services: *services,
	}
}
