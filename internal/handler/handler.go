package handler

import (
	"context"
	_ "github.com/NEKETSKY/mnemosyne/docs"
	"github.com/NEKETSKY/mnemosyne/internal/service"
	"github.com/NEKETSKY/mnemosyne/proto"
)

type Handler struct {
	ctx      context.Context
	services service.Service
	proto.GreeterServer
}

func NewHandler(ctx context.Context, services *service.Service) *Handler {
	return &Handler{
		ctx:      ctx,
		services: *services,
	}
}
